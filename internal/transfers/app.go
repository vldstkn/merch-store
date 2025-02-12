package transfers

import (
	"log/slog"
	"merch-store/internal/config"
	"merch-store/pkg/db"
)

type AppDeps struct {
	Config *config.Config
	DB     *db.DB
	Logger *slog.Logger
}

type App struct {
	Config *config.Config
	DB     *db.DB
	Logger *slog.Logger
}

func NewApp(deps *AppDeps) *App {
	return &App{
		Config: deps.Config,
		DB:     deps.DB,
		Logger: deps.Logger,
	}
}

func (app *App) Build() {

}
