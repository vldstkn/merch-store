package main

import (
	"log/slog"
	"merch-store/internal/config"
	"merch-store/internal/transfers"
	"merch-store/pkg/db"
	"merch-store/pkg/logger"
	"os"
)

func main() {
	conf := config.LoadConfig()
	database := db.NewDb(conf.DSN)
	log := logger.NewLogger(os.Stdout)
	app := transfers.NewApp(&transfers.AppDeps{
		DB:     database,
		Config: conf,
		Logger: log,
	})
	app.Logger.Info("Server start",
		slog.String("Name", "Products"),
		slog.String("Address", app.Config.TransfersAddress),
	)
}
