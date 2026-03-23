// internal/src/config.go

package src

import (
	"os"
	"fmt"
	"log"
	"strings"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type ConfigStructure struct{
	SERVER struct {
		Port int `yaml:"port"`
	} `yaml:"SERVER"`

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
}
