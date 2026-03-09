// internal/config/types.go

package handler

type ConfigSet struct {
	SERVER struct {
		Port int `yaml:"port"`
	} `yaml:"SERVER"`

	DB map[string]DBConfig `yaml:"DB"`
}

type DBConfig struct {
	Driver   string `yaml:"driver"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
}
