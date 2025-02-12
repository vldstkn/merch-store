package main

import (
	"merch-store/internal/config"
	"merch-store/internal/products"
	"merch-store/pkg/db"
	"merch-store/pkg/logger"
	"os"
)

func main() {
	conf := config.LoadConfig()
	database := db.NewDb(conf.DSN)
	log := logger.NewLogger(os.Stdout)
	app := products.NewApp(&products.AppDeps{
		DB:     database,
		Config: conf,
		Logger: log,
	})
	app.Run()
}
