package main

import (
	"merch_store/internal/config"
	"merch_store/internal/products"
	"merch_store/pkg/db"
	"merch_store/pkg/logger"
	"os"
)

func main() {
	mode := os.Getenv("APP_ENV")
	if mode == "" {
		mode = "dev"
	}
	conf := config.LoadConfig("./configs", mode)
	database := db.NewDb(conf.Database.Dsn)
	log := logger.NewLogger(os.Stdout)
	app := products.NewApp(&products.AppDeps{
		DB:     database,
		Config: conf,
		Logger: log,
		Mode:   mode,
	})
	app.Run()
}
