// internal/src/config_types.go

package src

type ConfigSet struct {
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
