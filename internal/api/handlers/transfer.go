package handlers

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"merch_store/internal/api/dto"
	"merch_store/internal/api/middleware"
	"merch_store/internal/config"
	http_errors "merch_store/internal/errors"
	grpc_conn "merch_store/pkg/grpc-conn"
	"merch_store/pkg/pb"
	"merch_store/pkg/req"
	"merch_store/pkg/res"
	"net/http"
)

type TransferHandlerDeps struct {
	Config *config.Config
	Logger *slog.Logger
}

type TransferHandler struct {
	Config          *config.Config
	Logger          *slog.Logger
	TransfersClient pb.TransfersClient
}

func NewTransferHandler(router chi.Router, deps *TransferHandlerDeps) error {

	transfersConn, err := grpc_conn.NewClientConn(deps.Config.Addresses.Transfers)
	if err != nil {
		return err
	}
	transfersClient := pb.NewTransfersClient(transfersConn)

	handler := TransferHandler{
		Config:          deps.Config,
		Logger:          deps.Logger,
		TransfersClient: transfersClient,
	}
	router.Group(func(r chi.Router) {
		r.Use(middleware.IsAuthed(handler.Config.Auth.Jwt))
		r.Post("/sendCoins", handler.SendCoins())
	})
	return nil
}

// SendCoins
//
//		@Summary		Отправить монеты другому пользователю.
//		@ID				sendCoins
//		@Produce		json
//		@Param			request	body	dto.SendCoinsReq		true	"Имя пользователя и сумма."
//		@Success		200		"Успешный ответ"
//		@Failure		400		{object}	dto.ErrorRes	"Неверный запрос."
//		@Failure		401		{object}	dto.ErrorRes	"Пользователь не авторизован."
//		@Failure		500		{object}	dto.ErrorRes	"Внутренняя ошибка сервера."
//		@Router			/sendCoins [post]
//	 @Security BearerAuth
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
		name := r.Context().Value("authData").(middleware.AuthData).Name

		fmt.Println(body, err, name)
		_, err = handler.TransfersClient.SendCoins(context.Background(), &pb.SendCoinsReq{
			FromUser: name,
			ToUser:   body.ToUser,
			Amount:   body.Amount,
		})
		if err != nil {
			mes, status := http_errors.HandleError(err)
			res.Json(w, dto.ErrorRes{
				Error: mes,
			}, status)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
