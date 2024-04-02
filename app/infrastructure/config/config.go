package config

type Config struct {
	Env   string `yaml:"env"`
	Host  string `yaml:"host"`
	Debug bool   `yaml:"debug"`
	Port  uint16 `yaml:"port"`
}
