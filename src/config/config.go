package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	DbName            string `yaml:"db_name"`
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	ApiKeyApiFootball string `yaml:"api_football_key"`
	UrlApiFootball    string `yaml:"api_football_url"`
}

func GetConfig() (Config, error) {
	configFile, err := os.ReadFile("./src/config/config.yaml")
	if err != nil {
		panic("cannot read configuration variables")
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
