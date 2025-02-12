package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"merch-store/internal/api/dto"
	"merch-store/internal/config"
	"merch-store/pkg/req"
	"merch-store/pkg/res"
	"net/http"
)

type TransferHandlerDeps struct {
	Config *config.Config
	Logger *slog.Logger
}

type TransferHandler struct {
	Config *config.Config
	Logger *slog.Logger
}

func NewTransferHandler(router *chi.Mux, deps *TransferHandlerDeps) error {
	handler := TransferHandler{
		Config: deps.Config,
		Logger: deps.Logger,
	}
	// TODO: проверка токена
	router.Post("/sendCoins", handler.SendCoins())
	return nil
}

func (handler *TransferHandler) SendCoins() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[dto.SendCoinsReq](r)
		if err != nil {
			handler.Logger.Error(err.Error(), slog.String("Error location", "req.HandleBody"))
			res.Json(w, dto.ErrorRes{
				Error: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		fmt.Println(body.Amount, body.ToUser)
	}
}
