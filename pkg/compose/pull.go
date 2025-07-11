/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package compose

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/compose-spec/compose-go/v2/types"
	"github.com/containerd/platforms"
	"github.com/distribution/reference"
	"github.com/docker/buildx/driver"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/registry"
	"github.com/hashicorp/go-multierror"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"xompose/pkg/api"
	"xompose/pkg/utils"
)

func (s *composeService) Pull(ctx context.Context, project *types.Project, options api.PullOptions) error {
	// return progress.RunWithTitle(ctx, func(ctx context.Context) error {
	return s.pull(ctx, project, options)
	// }, s.stdinfo(), "Pulling")
}

func (s *composeService) pull(ctx context.Context, project *types.Project, opts api.PullOptions) error { //nolint:gocyclo
	images, err := s.getLocalImagesDigests(ctx, project)
	if err != nil {
		return err
	}

	// w := progress.ContextWriter(ctx)
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(s.maxConcurrency)

	var (
		pullErrors        = make([]error, len(project.Services))
		imagesBeingPulled = map[string]string{}
	)

	i := 0
	// for name, service := range project.Services {
	for _, service := range project.Services {
		if service.Image == "" {
			// w.Event(progress.Event{
			// 	ID:     name,
			// 	Status: progress.Done,
			// 	Text:   "Skipped - No image to be pulled",
			// })
			continue
		}

		switch service.PullPolicy {
		case types.PullPolicyNever, types.PullPolicyBuild:
			// w.Event(progress.Event{
			// 	ID:     name,
			// 	Status: progress.Done,
			// 	Text:   "Skipped",
			// })
			continue
		case types.PullPolicyMissing, types.PullPolicyIfNotPresent:
			if imageAlreadyPresent(service.Image, images) {
				// w.Event(progress.Event{
				// 	ID:     name,
				// 	Status: progress.Done,
				// 	Text:   "Skipped - Image is already present locally",
				// })
				continue
			}
		}

		if service.Build != nil && opts.IgnoreBuildable {
			// w.Event(progress.Event{
			// 	ID:     name,
			// 	Status: progress.Done,
			// 	Text:   "Skipped - Image can be built",
			// })
			continue
		}

		// if s, ok := imagesBeingPulled[service.Image]; ok {
		if _, ok := imagesBeingPulled[service.Image]; ok {
			// w.Event(progress.Event{
			// 	ID:     name,
			// 	Status: progress.Done,
			// 	Text:   fmt.Sprintf("Skipped - Image is already being pulled by %v", s),
			// })
			continue
		}

		imagesBeingPulled[service.Image] = service.Name

		idx := i
		eg.Go(func() error {
			// _, err := s.pullServiceImage(ctx, service, LoadDefaultConfigFile(), w, opts.Quiet, project.Environment["DOCKER_DEFAULT_PLATFORM"])
			// _, err := s.pullServiceImage(ctx, service, LoadDefaultConfigFile(), w, opts.Quiet, project.Environment["DOCKER_DEFAULT_PLATFORM"])
			_, err := s.pullServiceImage(ctx, service, LoadDefaultConfigFile(), opts.Quiet, project.Environment["DOCKER_DEFAULT_PLATFORM"])
			if err != nil {
				pullErrors[idx] = err
				// if service.Build != nil {
				// 	mustBuild = append(mustBuild, service.Name)
				// }
				if !opts.IgnoreFailures && service.Build == nil {
					if s.dryRun {
						// w.Event(progress.Event{
						// 	ID:     name,
						// 	Status: progress.Error,
						// 	Text:   fmt.Sprintf(" - Pull error for image: %s", service.Image),
						// })
					}
					// fail fast if image can't be pulled nor built
					return err
				}
			}
			return nil
		})
		i++
	}

	if err = eg.Wait(); err != nil {
		return err
	}
	if opts.IgnoreFailures {
		return nil
	}
	return multierror.Append(nil, pullErrors...).ErrorOrNil()
}

func imageAlreadyPresent(serviceImage string, localImages map[string]api.ImageSummary) bool {
	normalizedImage, err := reference.ParseDockerRef(serviceImage)
	if err != nil {
		return false
	}
	tagged, ok := normalizedImage.(reference.NamedTagged)
	if !ok {
		return false
	}
	_, ok = localImages[serviceImage]
	return ok && tagged.Tag() != "latest"
}

func getUnwrappedErrorMessage(err error) string {
	derr := errors.Unwrap(err)
	if derr != nil {
		return getUnwrappedErrorMessage(derr)
	}
	return err.Error()
}

func (s *composeService) pullServiceImage(ctx context.Context, service types.ServiceConfig,
	// configFile driver.Auth, w progress.Writer, quietPull bool, defaultPlatform string,
	configFile driver.Auth, quietPull bool, defaultPlatform string,
) (string, error) {
	// w.Event(progress.Event{
	// 	ID:     service.Name,
	// 	Status: progress.Working,
	// 	Text:   "Pulling",
	// })
	ref, err := reference.ParseNormalizedNamed(service.Image)
	if err != nil {
		return "", err
	}

	encodedAuth, err := encodedAuth(ref, configFile)
	if err != nil {
		return "", err
	}

	platform := service.Platform
	if platform == "" {
		platform = defaultPlatform
	}

	stream, err := s.apiClient().ImagePull(ctx, service.Image, image.PullOptions{
		RegistryAuth: encodedAuth,
		Platform:     platform,
	})

	if ctx.Err() != nil {
		// w.Event(progress.Event{
		// 	ID:         service.Name,
		// 	Status:     progress.Warning,
		// 	StatusText: "Interrupted",
		// })
		return "", nil
	}

	// check if has error and the service has a build section
	// then the status should be warning instead of error
	if err != nil && service.Build != nil {
		// w.Event(progress.Event{
		// 	ID:         service.Name,
		// 	Status:     progress.Warning,
		// 	Text:       "Warning",
		// 	StatusText: getUnwrappedErrorMessage(err),
		// })
		return "", err
	}

	if err != nil {
		// w.Event(progress.Event{
		// 	ID:         service.Name,
		// 	Status:     progress.Error,
		// 	Text:       "Error",
		// 	StatusText: getUnwrappedErrorMessage(err),
		// })
		return "", err
	}

	dec := json.NewDecoder(stream)
	for {
		var jm jsonmessage.JSONMessage
		if err := dec.Decode(&jm); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", err
		}
		if jm.Error != nil {
			return "", errors.New(jm.Error.Message)
		}
		// if !quietPull {
		// 	toPullProgressEvent(service.Name, jm, w)
		// }
	}
	// w.Event(progress.Event{
	// 	ID:     service.Name,
	// 	Status: progress.Done,
	// 	Text:   "Pulled",
	// })

	inspected, err := s.apiClient().ImageInspect(ctx, service.Image)
	if err != nil {
		return "", err
	}
	return inspected.ID, nil
}

func encodedAuth(ref reference.Named, configFile driver.Auth) (string, error) {
	repoInfo, err := registry.ParseRepositoryInfo(ref)
	if err != nil {
		return "", err
	}

	key := registry.GetAuthConfigKey(repoInfo.Index)
	authConfig, err := configFile.GetAuthConfig(key)
	if err != nil {
		return "", err
	}

	buf, err := json.Marshal(authConfig)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buf), nil
}

func (s *composeService) pullRequiredImages(ctx context.Context, project *types.Project, images map[string]api.ImageSummary, quietPull bool) error {
	needPull := map[string]types.ServiceConfig{}
	for name, service := range project.Services {
		pull, err := mustPull(service, images)
		if err != nil {
			return err
		}
		if pull {
			needPull[name] = service
		}
		for i, vol := range service.Volumes {
			if vol.Type == types.VolumeTypeImage {
				if _, ok := images[vol.Source]; !ok {
					// Hack: create a fake ServiceConfig so we pull missing volume image
					n := fmt.Sprintf("%s:volume %d", name, i)
					needPull[n] = types.ServiceConfig{
						Name:  n,
						Image: vol.Source,
					}
				}
			}
		}

	}
	if len(needPull) == 0 {
		return nil
	}

	// return progress.Run(ctx, func(ctx context.Context) error {
	// w := progress.ContextWriter(ctx)
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(s.maxConcurrency)
	pulledImages := map[string]api.ImageSummary{}
	var mutex sync.Mutex
	for name, service := range needPull {
		eg.Go(func() error {
			// id, err := s.pullServiceImage(ctx, service, s.configFile(), w, quietPull, project.Environment["DOCKER_DEFAULT_PLATFORM"])
			// id, err := s.pullServiceImage(ctx, service, LoadDefaultConfigFile(), w, quietPull, project.Environment["DOCKER_DEFAULT_PLATFORM"])
			id, err := s.pullServiceImage(ctx, service, LoadDefaultConfigFile(), quietPull, project.Environment["DOCKER_DEFAULT_PLATFORM"])
			mutex.Lock()
			defer mutex.Unlock()
			pulledImages[name] = api.ImageSummary{
				ID:          id,
				Repository:  service.Image,
				LastTagTime: time.Now(),
			}
			// if err != nil && isServiceImageToBuild(service, project.Services) {
			// 	// image can be built, so we can ignore pull failure
			// 	return nil
			// }
			return err
		})
	}
	err := eg.Wait()
	for i, service := range needPull {
		if pulledImages[i].ID != "" {
			images[service.Image] = pulledImages[i]
		}
	}
	return err
	// }, s.stdinfo())
}

func mustPull(service types.ServiceConfig, images map[string]api.ImageSummary) (bool, error) {
	if service.Provider != nil {
		return false, nil
	}
	if service.Image == "" {
		return false, nil
	}
	policy, duration, err := service.GetPullPolicy()
	if err != nil {
		return false, err
	}
	switch policy {
	case types.PullPolicyAlways:
		// force pull
		return true, nil
	case types.PullPolicyNever, types.PullPolicyBuild:
		return false, nil
	case types.PullPolicyRefresh:
		img, ok := images[service.Image]
		if !ok {
			return true, nil
		}
		return time.Now().After(img.LastTagTime.Add(duration)), nil
	default: // Pull if missing
		_, ok := images[service.Image]
		return !ok, nil
	}
}

// func isServiceImageToBuild(service types.ServiceConfig, services types.Services) bool {
// 	if service.Build != nil {
// 		return true
// 	}

// 	if service.Image == "" {
// 		// N.B. this should be impossible as service must have either `build` or `image` (or both)
// 		return false
// 	}

// 	// look through the other services to see if another has a build definition for the same
// 	// image name
// 	for _, svc := range services {
// 		if svc.Image == service.Image && svc.Build != nil {
// 			return true
// 		}
// 	}
// 	return false
// }

const (
	PreparingPhase         = "Preparing"
	WaitingPhase           = "Waiting"
	PullingFsPhase         = "Pulling fs layer"
	DownloadingPhase       = "Downloading"
	DownloadCompletePhase  = "Download complete"
	ExtractingPhase        = "Extracting"
	VerifyingChecksumPhase = "Verifying Checksum"
	AlreadyExistsPhase     = "Already exists"
	PullCompletePhase      = "Pull complete"
)

// func toPullProgressEvent(parent string, jm jsonmessage.JSONMessage, w progress.Writer) {
// 	if jm.ID == "" || jm.Progress == nil {
// 		return
// 	}

// 	var (
// 		text    string
// 		total   int64
// 		percent int
// 		current int64
// 		status  = progress.Working
// 	)

// 	text = jm.Progress.String()

// 	switch jm.Status {
// 	case PreparingPhase, WaitingPhase, PullingFsPhase:
// 		percent = 0
// 	case DownloadingPhase, ExtractingPhase, VerifyingChecksumPhase:
// 		if jm.Progress != nil {
// 			current = jm.Progress.Current
// 			total = jm.Progress.Total
// 			if jm.Progress.Total > 0 {
// 				percent = int(jm.Progress.Current * 100 / jm.Progress.Total)
// 			}
// 		}
// 	case DownloadCompletePhase, AlreadyExistsPhase, PullCompletePhase:
// 		status = progress.Done
// 		percent = 100
// 	}

// 	if strings.Contains(jm.Status, "Image is up to date") ||
// 		strings.Contains(jm.Status, "Downloaded newer image") {
// 		status = progress.Done
// 		percent = 100
// 	}

// 	if jm.Error != nil {
// 		status = progress.Error
// 		text = jm.Error.Message
// 	}

// 	w.Event(progress.Event{
// 		ID:         jm.ID,
// 		ParentID:   parent,
// 		Current:    current,
// 		Total:      total,
// 		Percent:    percent,
// 		Text:       jm.Status,
// 		Status:     status,
// 		StatusText: text,
// 	})
// }

func (s *composeService) getLocalImagesDigests(ctx context.Context, project *types.Project) (map[string]api.ImageSummary, error) {
	imageNames := utils.Set[string]{}
	for _, s := range project.Services {
		imageNames.Add(api.GetImageNameOrDefault(s, project.Name))
		for _, volume := range s.Volumes {
			if volume.Type == types.VolumeTypeImage {
				imageNames.Add(volume.Source)
			}
		}
	}
	images, err := s.getImageSummaries(ctx, imageNames.Elements())
	if err != nil {
		return nil, err
	}

	for i, service := range project.Services {
		imgName := api.GetImageNameOrDefault(service, project.Name)
		img, ok := images[imgName]
		if !ok {
			continue
		}
		if service.Platform != "" {
			platform, err := platforms.Parse(service.Platform)
			if err != nil {
				return nil, err
			}
			inspect, err := s.apiClient().ImageInspect(ctx, img.ID)
			if err != nil {
				return nil, err
			}
			actual := specs.Platform{
				Architecture: inspect.Architecture,
				OS:           inspect.Os,
				Variant:      inspect.Variant,
			}
			if !platforms.NewMatcher(platform).Match(actual) {
				logrus.Debugf("local image %s doesn't match expected platform %s", service.Image, service.Platform)
				// there is a local image, but it's for the wrong platform, so
				// pretend it doesn't exist so that we can pull/build an image
				// for the correct platform instead
				delete(images, imgName)
			}
		}

		project.Services[i].CustomLabels.Add(api.ImageDigestLabel, img.ID)

	}

	return images, nil
}
