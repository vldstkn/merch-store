package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"merch_store/internal/api/dto"
	"merch_store/internal/api/mappers"
	"merch_store/internal/api/middleware"
	"merch_store/internal/config"
	http_errors "merch_store/internal/errors"
	grpc_conn "merch_store/pkg/grpc-conn"
	"merch_store/pkg/pb"
	"merch_store/pkg/req"
	"merch_store/pkg/res"
	"net/http"
)

type AccountHandlerDeps struct {
	Config *config.Config
	Logger *slog.Logger
}

type AccountHandler struct {
	Config        *config.Config
	Logger        *slog.Logger
	AccountClient pb.AccountClient
}

func NewAccountHandler(router chi.Router, deps *AccountHandlerDeps) error {
	accountConn, err := grpc_conn.NewClientConn(deps.Config.Addresses.Account)
	if err != nil {
		return err
	}
	accountClient := pb.NewAccountClient(accountConn)
	handler := AccountHandler{
		Config:        deps.Config,
		Logger:        deps.Logger,
		AccountClient: accountClient,
	}
	router.Post("/auth", handler.Auth())
	router.Route("/info", func(r chi.Router) {
		r.Use(middleware.IsAuthed(handler.Config.Auth.Jwt))
		r.Get("/", handler.Info())
	})
	return nil
}

// Auth
//
//		@Summary		Аутентификация и получение JWT-токена.
//		@Description	При первой аутентификации пользователь создается автоматически.
//		@ID				auth
//		@Accept			json
//		@Produce		json
//	  @Param			request	body	dto.AuthReq		true	"Имя пользователя и пароль."
//		@Success		201		"Успешный ответ"
//		@Failure		400		{object}	dto.ErrorRes	"Неверный запрос."
//		@Failure		500		{object}	dto.ErrorRes	"Внутренняя ошибка сервера."
//		@Router			/auth [post]
func (handler *AccountHandler) Auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[dto.AuthReq](r)
		if err != nil {
			handler.Logger.Error(err.Error(), slog.String("Error location", "req.HandleBody"))
			res.Json(w, dto.ErrorRes{
				Error: http.StatusText(http.StatusBadRequest),
			}, http.StatusBadRequest)
			return
		}
		response, err := handler.AccountClient.Auth(context.Background(), &pb.AuthReq{
			Password: body.Password,
			Username: body.Username,
		})
		if err != nil {
			mes, code := http_errors.HandleError(err)
			res.Json(w, dto.ErrorRes{
				Error: mes,
			}, code)
			return
		}
		res.Json(w, response, http.StatusOK)
	}
}

// Info
//
//		@Summary		Получить информацию о пользователе.
//		@Description	Получить информацию о балансе, купленных предметах, истории переводов.
//		@ID				info
//		@Produce		json
//		@Success		200 {object} dto.GetInfoRes		"Успешный ответ"
//		@Failure		400		{object}	dto.ErrorRes	"Неверный запрос."
//		@Failure		401		{object}	dto.ErrorRes	"Пользователь не авторизован."
//		@Failure		500		{object}	dto.ErrorRes	"Внутренняя ошибка сервера."
//		@Router			/info [get]
//	 @Security BearerAuth
func (handler *AccountHandler) Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.Context().Value("authData").(middleware.AuthData).Name
		response, err := handler.AccountClient.GetInfo(context.Background(), &pb.GetInfoReq{
			UserName: name,
		})
		if err != nil {
			res.Json(w, dto.ErrorRes{
				Error: http.StatusText(http.StatusInternalServerError),
			}, http.StatusInternalServerError)
			return
		}
		res.Json(w, dto.GetInfoRes{
			Inventory:    mappers.InventoryRepFromGrpcToDto(response.Inventory),
			Coins:        response.Coins,
			CoinsHistory: mappers.CoinsHistoryFromGrpcToDto(response.CoinsHistory),
		}, http.StatusOK)
	}
}
