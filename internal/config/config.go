package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
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
}

func LoadConfig(path, mode string) *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	viper.SetConfigName("config." + env)
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
