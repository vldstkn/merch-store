package api

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"merch_store/internal/api/handlers"
	"merch_store/internal/config"
)

type HandlersDeps struct {
	Config *config.Config
	Logger *slog.Logger
}

func NewHandlers(router chi.Router, deps *HandlersDeps) error {
	err := handlers.NewAccountHandler(router, &handlers.AccountHandlerDeps{
		Config: deps.Config,
		Logger: deps.Logger,
	})
	if err != nil {
		return err
	}
	err = handlers.NewProductsHandler(router, &handlers.ProductsHandlerDeps{
		Config: deps.Config,
		Logger: deps.Logger,
	})
	if err != nil {
		return err
	}
	err = handlers.NewTransferHandler(router, &handlers.TransferHandlerDeps{
		Logger: deps.Logger,
		Config: deps.Config,
	})
	if err != nil {
		return err
	}
	return nil
}
