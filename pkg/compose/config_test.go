package compose

import (
	"testing"
)

func TestLoadDefaultConfigFile(t *testing.T) {
	cfg := LoadDefaultConfigFile()
	t.Log(cfg)

	// c, err := cfg.GetAuthConfig("https://index.docker.io/v1/")
	c, err := cfg.GetAuthConfig("registry.cn-beijing.aliyuncs.com")
	if err != nil {
		t.Log(err)
	}
	t.Log(c)
}
