package main

import (
	"merch-store/internal/account"
	"merch-store/internal/config"
	"merch-store/pkg/db"
	"merch-store/pkg/logger"
	"os"
)

func main() {
	conf := config.LoadConfig()
	database := db.NewDb(conf.DSN)
	log := logger.NewLogger(os.Stdout)
	app := account.NewApp(&account.AppDeps{
		DB:     database,
		Config: conf,
		Logger: log,
	})
	app.Run()
}
