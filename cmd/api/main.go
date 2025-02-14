package main

import (
	"context"
	"log/slog"
	_ "merch_store/docs"
	"merch_store/internal/api"
	"merch_store/internal/config"
	"merch_store/pkg/logger"
	"net/http"
	"os"
)

//	@title			Swagger Example API
//	@contact.email	vldstkn.develop@gmail.com
//  @Version 1.0

//	@host		localhost:8080
//	@BasePath	/api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	mode := os.Getenv("APP_ENV")
	if mode == "" {
		mode = "dev"
	}
	conf := config.LoadConfig("./configs", mode)
	log := logger.NewLogger(os.Stdout)
	app := api.NewApp(&api.AppDeps{
		Config: conf,
		Logger: log,
		Mode:   mode,
	})
	router, err := app.Build()
	if err != nil {
		return
	}
	server := http.Server{
		Addr:    app.Config.Addresses.Api,
		Handler: router,
	}
	defer server.Shutdown(context.Background())
	app.Logger.Info("Server starts",
		slog.String("Address", app.Config.Addresses.Api),
		slog.String("Name", "Api"),
		slog.String("Mode", mode),
	)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Error(err.Error(),
			slog.String("Api address", app.Config.Addresses.Api),
		)
	}
}
