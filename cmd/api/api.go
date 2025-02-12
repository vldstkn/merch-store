package main

import (
	"log/slog"
	"merch-store/internal/api"
	"merch-store/internal/config"
	"merch-store/pkg/logger"
	"net/http"
	"os"
)

func main() {
	conf := config.LoadConfig()
	log := logger.NewLogger(os.Stdout)
	app := api.NewApp(&api.AppDeps{
		Config: conf,
		Logger: log,
	})
	router, err := app.Build()
	if err != nil {
		return
	}
	server := http.Server{
		Addr:    app.Config.ApiAddress,
		Handler: router,
	}
	app.Logger.Info("Server starts",
		slog.String("Address", app.Config.ApiAddress),
		slog.String("Name", "Api"),
	)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Error(err.Error(),
			slog.String("Api address", app.Config.ApiAddress),
		)
	}
	server.Close()
}
