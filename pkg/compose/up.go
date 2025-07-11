package compose

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"xompose/pkg/api"

	"github.com/compose-spec/compose-go/v2/types"
	"github.com/containerd/errdefs"
	"github.com/docker/cli/cli"
	"github.com/eiannone/keyboard"
	"github.com/hashicorp/go-multierror"
)

func (s *composeService) Up(ctx context.Context, project *types.Project, options api.UpOptions) error { //nolint:gocyclo
	// err := progress.Run(ctx, tracing.SpanWrapFunc("project/up", tracing.ProjectOptions(ctx, project), func(ctx context.Context) error {
	err := s.create(ctx, project, options.Create)
	if err != nil {
		return err
	}
	if options.Start.Attach == nil {
		return s.start(ctx, project.Name, options.Start, nil)
	}
	// return nil
	// }), s.stdinfo())
	// if err != nil {
	// 	return err
	// }

	if options.Start.Attach == nil {
		return err
	}
	if s.dryRun {
		// _, _ = fmt.Fprintln(s.stdout(), "end of 'compose up' output, interactive run is not supported in dry-run mode")
		return err
	}

	var eg multierror.Group

	// if we get a second signal during shutdown, we kill the services
	// immediately, so the channel needs to have sufficient capacity or
	// we might miss a signal while setting up the second channel read
	// (this is also why signal.Notify is used vs signal.NotifyContext)
	signalChan := make(chan os.Signal, 2)
	defer close(signalChan)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signalChan)
	var isTerminated atomic.Bool
	printer := newLogPrinter(options.Start.Attach)

	// watcher, err := NewWatcher(project, options, s.watch)
	// if err != nil && options.Start.Watch {
	// 	return err
	// }

	// var navigationMenu *formatter.LogKeyboard
	// var kEvents <-chan keyboard.KeyEvent
	// if options.Start.NavigationMenu {
	// 	kEvents, err = keyboard.GetKeys(100)
	// 	if err != nil {
	// 		logrus.Warnf("could not start menu, an error occurred while starting: %v", err)
	// 		options.Start.NavigationMenu = false
	// 	} else {
	// 		defer keyboard.Close() //nolint:errcheck
	// 		// isDockerDesktopActive := s.isDesktopIntegrationActive()
	// 		tracing.KeyboardMetrics(ctx, options.Start.NavigationMenu, watcher != nil)
	// 		navigationMenu = formatter.NewKeyboardManager(false, signalChan, options.Start.Watch, watcher)
	// 	}
	// }

	doneCh := make(chan bool)
	eg.Go(func() error {
		first := true
		gracefulTeardown := func() {
			printer.Cancel()
			// _, _ = fmt.Fprintln(s.stdinfo(), "Gracefully stopping... (press Ctrl+C again to force)")
			eg.Go(func() error {
				err := s.Stop(context.WithoutCancel(ctx), project.Name, api.StopOptions{
					Services: options.Create.Services,
					Project:  project,
				})
				isTerminated.Store(true)
				return err
			})
			first = false
		}

		for {
			select {
			case <-doneCh:
				// if watcher != nil {
				// 	return watcher.Stop()
				// }
				return nil
			case <-ctx.Done():
				if first {
					gracefulTeardown()
				}
			case <-signalChan:
				if first {
					keyboard.Close() //nolint:errcheck
					gracefulTeardown()
					break
				}
				eg.Go(func() error {
					err := s.kill(context.WithoutCancel(ctx), project.Name, api.KillOptions{
						Services: options.Create.Services,
						Project:  project,
						All:      true,
					})
					// Ignore errors indicating that some of the containers were already stopped or removed.
					if errdefs.IsNotFound(err) || errdefs.IsConflict(err) {
						return nil
					}

					return err
				})
				return nil
				// case event := <-kEvents:
				// navigationMenu.HandleKeyEvents(ctx, event, project, options)
			}
		}
	})

	var exitCode int
	eg.Go(func() error {
		code, err := printer.Run(options.Start.OnExit, options.Start.ExitCodeFrom, func() error {
			// _, _ = fmt.Fprintln(s.stdinfo(), "Aborting on container exit...")
			// return progress.Run(ctx, func(ctx context.Context) error {
			// 	return s.Stop(ctx, project.Name, api.StopOptions{
			// 		Services: options.Create.Services,
			// 		Project:  project,
			// 	})
			// }, s.stdinfo())
			return s.Stop(ctx, project.Name, api.StopOptions{
				Services: options.Create.Services,
				Project:  project,
			})
		})
		exitCode = code
		return err
	})

	// if options.Start.Watch && watcher != nil {
	// 	err = watcher.Start(ctx)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// We use the parent context without cancellation as we manage sigterm to stop the stack
	err = s.start(context.WithoutCancel(ctx), project.Name, options.Start, printer.HandleEvent)
	if err != nil && !isTerminated.Load() { // Ignore error if the process is terminated
		return err
	}

	// Signal for the signal-handler goroutines to stop
	close(doneCh)

	printer.Stop()

	err = eg.Wait().ErrorOrNil()
	if exitCode != 0 {
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		}
		return cli.StatusError{StatusCode: exitCode, Status: errMsg}
	}
	return err
}
