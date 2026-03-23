// internal/src/config.go

package src

import (
	"os"
<<<<<<< HEAD
	"fmt"
	"log"
=======
	"strings"
	"os/exec"
>>>>>>> 1d214e3a2fe5e4da03403606d81c52ca4fa2af06

	"gopkg.in/yaml.v3"
)

type ConfigStructure struct{
	SERVER struct {
		Port int `yaml:"port"`
	} `yaml:"SERVER"`

<<<<<<< HEAD
	DB            map[string]DBConfig `yaml:"DB"`
	APP_STRUCTURE map[string]string   `yaml:"APP_STRUCTURE"`
}

type DBConfig struct {
	Driver   string `yaml:"driver"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

var (
	config ConfigStructure
	cnfLoaded bool
)

func ConfigGet() ConfigStructure {
	if !cnfLoaded {
		configLoad()
	}
	
	return config
}

func configLoad() {
	cnfLoaded 	= true
	config 		= ConfigStructure{}
	path 		:= DirBase() + "/internal/config/config.yaml"
	
	data, err := os.ReadFile(path)
	if err != nil {
		msg := fmt.Sprintf("[config.go] Não foi possível abrir o arquivo [%s]", path)
		log.Fatal(msg)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		msg := fmt.Sprintf("[config.go] Não foi possível parsear as configurações do arquivo [%s]", path)
		log.Fatal(msg)
	}
=======
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
>>>>>>> 1d214e3a2fe5e4da03403606d81c52ca4fa2af06
}

func CurrentDir() string {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(out))
}
