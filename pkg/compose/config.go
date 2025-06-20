package compose

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/cli/cli/config/types"
	"github.com/docker/docker-credential-helpers/client"
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/pkg/errors"
)

// ConfigFile ~/.docker/config.json file info
type ConfigFile struct {
	Auths            map[string]types.AuthConfig `json:"auths"`
	CredentialsStore string                      `json:"credsStore,omitempty"`
	CurrentContext   string                      `json:"currentContext,omitempty"`
	Filename         string                      `json:"-"` // Note: for internal use only
}

func New(fn string) *ConfigFile {
	return &ConfigFile{
		Auths:    make(map[string]types.AuthConfig),
		Filename: fn,
	}
}

func (config *ConfigFile) GetAuthConfigs() map[string]types.AuthConfig {
	if config.Auths == nil {
		config.Auths = make(map[string]types.AuthConfig)
	}
	return config.Auths
}

// GetCredentialsStore returns a new credentials store from the settings in the
// configuration file
// func (configFile *ConfigFile) GetCredentialsStore(registryHostname string) credentials.Store {
// 	if helper := getConfiguredCredentialStore(configFile, registryHostname); helper != "" {
// 		return newNativeStore(configFile, helper)
// 	}
// 	return credentials.NewFileStore(configFile)
// }

// getConfiguredCredentialStore returns the credential helper configured for the
// given registry, the default credsStore, or the empty string if neither are
// configured.
// func getConfiguredCredentialStore(c *ConfigFile, registryHostname string) string {
// 	if c.CredentialHelpers != nil && registryHostname != "" {
// 		if helper, exists := c.CredentialHelpers[registryHostname]; exists {
// 			return helper
// 		}
// 	}
// 	return c.CredentialsStore
// }

func (configFile *ConfigFile) GetAuthConfig(registryHostname string) (types.AuthConfig, error) {

	authConfig, _ := configFile.GetAuthConfigs()[registryHostname]
	if configFile.CredentialsStore != "" {
		creds, err := client.Get(client.NewShellProgramFunc("docker-credential-"+configFile.CredentialsStore), registryHostname)
		if err != nil {
			if credentials.IsErrCredentialsNotFound(err) {
				// do not return an error if the credentials are not
				// in the keychain. Let docker ask for new credentials.
				return types.AuthConfig{}, nil
			}
			return types.AuthConfig{}, err
		}
		if creds.Username == "<token>" {
			authConfig.IdentityToken = creds.Secret
		} else {
			authConfig.Username = creds.Username
			authConfig.Password = creds.Secret
		}
		authConfig.ServerAddress = registryHostname
	}
	return authConfig, nil
}

// decodeAuth decodes a base64 encoded string and returns username and password
func decodeAuth(authStr string) (string, string, error) {
	if authStr == "" {
		return "", "", nil
	}

	decLen := base64.StdEncoding.DecodedLen(len(authStr))
	decoded := make([]byte, decLen)
	authByte := []byte(authStr)
	n, err := base64.StdEncoding.Decode(decoded, authByte)
	if err != nil {
		return "", "", err
	}
	if n > decLen {
		return "", "", errors.Errorf("Something went wrong decoding auth config")
	}
	userName, password, ok := strings.Cut(string(decoded), ":")
	if !ok || userName == "" {
		return "", "", errors.Errorf("Invalid auth configuration file")
	}
	return userName, strings.Trim(password, "\x00"), nil
}

func (configFile *ConfigFile) LoadFromReader(configData io.Reader) error {
	if err := json.NewDecoder(configData).Decode(configFile); err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	var err error
	for addr, ac := range configFile.Auths {
		if ac.Auth != "" {
			ac.Username, ac.Password, err = decodeAuth(ac.Auth)
			if err != nil {
				return err
			}
		}
		ac.Auth = ""
		ac.ServerAddress = addr
		configFile.Auths[addr] = ac
	}
	return nil
}

func LoadDefaultConfigFile() *ConfigFile {
	home, _ := os.UserHomeDir()
	filename := filepath.Join(home, ".docker", "config.json")
	configFile := New(filename)
	file, err := os.Open(filename)
	if err != nil {
		return configFile
	}
	defer file.Close()
	err = configFile.LoadFromReader(file)
	if err != nil {
		err = errors.Wrapf(err, "parsing config file (%s)", filename)
	}

	return configFile
}
