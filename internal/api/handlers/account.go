package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"merch-store/internal/api/dto"
	"merch-store/internal/api/middleware"
	"merch-store/internal/config"
	http_errors "merch-store/internal/errors"
	grpc_conn "merch-store/pkg/grpc-conn"
	pb_account "merch-store/pkg/pb"
	"merch-store/pkg/req"
	"merch-store/pkg/res"
	"net/http"
)

type AccountHandlerDeps struct {
	Config *config.Config
	Logger *slog.Logger
}

type AccountHandler struct {
	Config        *config.Config
	Logger        *slog.Logger
	AccountClient pb_account.AccountClient
}

func NewAccountHandler(router *chi.Mux, deps *AccountHandlerDeps) error {
	accountConn, err := grpc_conn.NewClientConn(deps.Config.AccountAddress)
	if err != nil {
		return err
	}
	accountClient := pb_account.NewAccountClient(accountConn)
	handler := AccountHandler{
		Config:        deps.Config,
		Logger:        deps.Logger,
		AccountClient: accountClient,
	}
	router.Post("/auth", handler.Auth())
	router.Route("/info", func(r chi.Router) {
		r.Use(middleware.IsAuthed(handler.Config.JWTSecret))
		r.Get("/", handler.Info())
	})
	return nil
}

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
		response, err := handler.AccountClient.Auth(context.Background(), &pb_account.AuthReq{
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
func (handler *AccountHandler) Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("authData").(middleware.AuthData).Id
		response, err := handler.AccountClient.GetInfo(context.Background(), &pb_account.GetInfoReq{
			UserId: id,
		})
		if err != nil {
			res.Json(w, dto.ErrorRes{
				Error: http.StatusText(http.StatusInternalServerError),
			}, http.StatusInternalServerError)
			return
		}
		res.Json(w, response, http.StatusOK)
	}
}
