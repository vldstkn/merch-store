package api

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"merch_store/internal/config"
)

type AppDeps struct {
	Config *config.Config
	Logger *slog.Logger
	Mode   string
}

type App struct {
	Config *config.Config
	Logger *slog.Logger
	Mode   string
}

func NewApp(deps *AppDeps) *App {
	return &App{
		Config: deps.Config,
		Logger: deps.Logger,
		Mode:   deps.Mode,
	}
}

func (app *App) Build() (*chi.Mux, error) {
	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		NewHandlers(r, &HandlersDeps{
			Config: app.Config,
			Logger: app.Logger,
		})
	})

	return router, nil
}
