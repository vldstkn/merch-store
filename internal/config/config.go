package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Addresses struct {
		Api       string `yaml:"api"`
		Account   string `yaml:"account"`
		Products  string `yaml:"products"`
		Transfers string `yaml:"transfers"`
	} `yaml:"addresses"`
	Database struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"database"`
	Auth struct {
		Jwt string `yaml:"jwt"`
	} `yaml:"auth"`
	Public struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"database"`
	} `yaml:"public"`
}

func LoadConfig(path, mode string) *Config {
	viper.SetConfigName("config." + mode)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("bad path to config: %v", path)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("YAML parsing error")
	}
	return &config
}
