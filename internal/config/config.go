package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ApiAddress       string
	AccountAddress   string
	ProductsAddress  string
	TransfersAddress string
	DSN              string
	JWTSecret        string
}

func LoadConfig() *Config {
	err := godotenv.Load("./configs/.env")
	if err != nil {
		panic(err)
	}
	return &Config{
		ApiAddress:       os.Getenv("ApiAddress"),
		AccountAddress:   os.Getenv("AccountAddress"),
		ProductsAddress:  os.Getenv("ProductsAddress"),
		TransfersAddress: os.Getenv("TransfersAddress"),
		DSN:              os.Getenv("DSN"),
		JWTSecret:        os.Getenv("JWTSecret"),
	}
}
