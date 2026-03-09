// internal/config/config.go

package handler

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigHandler struct{}

var Config ConfigHandler = ConfigHandler{}

var cfg *ConfigSet

func (Config *ConfigHandler) Load() *Config {
	if cfg != nil {
		return cfg
	}

	cfg = &ConfigSet{}
	contexts := [...]string{
		"db",
		"server",
	}

	for _, ctx := range contexts {
		path := fmt.Sprintf("config/%s.yaml", ctx)
		data, err := os.ReadFile(path)

		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(data, cfg)

		if err != nil {
			panic(err)
		}
	}

	return cfg
}
