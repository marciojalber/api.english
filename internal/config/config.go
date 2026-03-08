// internal/config/config.go

package config

import (
	"os"
	"fmt"

	"gopkg.in/yaml.v3"
)

var cfg *Config

func Load() *Config {
	if cfg != nil {
		return cfg
	}

	cfg 			= &Config{}
	contexts 		:= [...]string{
		"db",
		"server",
	}

	for _, ctx := range contexts {
		path 		:= fmt.Sprintf("config/%s.yaml", ctx)
		data, err 	:= os.ReadFile(path)
		
		if err != nil {
			panic(err)
		}

		err 		= yaml.Unmarshal(data, cfg)
		
		if err != nil {
			panic(err)
		}
	}

	return cfg
}
