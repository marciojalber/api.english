// internal/src/config.go

package src

import (
	"fmt"
	"os"
	"strings"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type ConfigHandler struct{}

var cfg *ConfigSet
var Config ConfigHandler = ConfigHandler{}

func (config *ConfigHandler) Load() *ConfigSet {
	if cfg != nil {
		return cfg
	}

	cfg = &ConfigSet{}
	contexts := [...]string{
		"db",
		"server",
		"system",
	}

	for _, ctx := range contexts {
		path := fmt.Sprintf(CurrentDir() + "/config/%s.yaml", ctx)
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

func CurrentDir() string {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(out))
}
