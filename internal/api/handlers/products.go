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
	"merch-store/pkg/pb"
	"merch-store/pkg/res"
	"net/http"
	"strconv"
)

type ProductsHandlerDeps struct {
	Config *config.Config
	Logger *slog.Logger
}

type ProductsHandler struct {
	Config         *config.Config
	Logger         *slog.Logger
	ProductsClient pb.ProductClient
}

func NewProductsHandler(router *chi.Mux, deps *ProductsHandlerDeps) error {
	productsConn, err := grpc_conn.NewClientConn(deps.Config.ProductsAddress)
	if err != nil {
		return err
	}
	productsClient := pb.NewProductClient(productsConn)
	handler := ProductsHandler{
		Config:         deps.Config,
		Logger:         deps.Logger,
		ProductsClient: productsClient,
	}
	router.Group(func(r chi.Router) {
		r.Use(middleware.IsAuthed(handler.Config.JWTSecret))
		r.Get("/buy/{item}", handler.Buy())
	})
	return nil
}

func (handler *ProductsHandler) Buy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productStr := chi.URLParam(r, "item")
		productId, err := strconv.Atoi(productStr)
		if err != nil {
			handler.Logger.Error(err.Error(),
				slog.String("Error location", "strconv.Atoi"),
				slog.String("Item id", productStr),
			)
			res.Json(w, dto.ErrorRes{
				Error: http.StatusText(http.StatusBadRequest),
			}, http.StatusBadRequest)
			return
		}
		userId := r.Context().Value("authData").(middleware.AuthData).Id
		_, err = handler.ProductsClient.Buy(context.Background(), &pb.BuyReq{
			UserId:    userId,
			ProductId: int64(productId),
		})
		if err != nil {
			mes, code := http_errors.HandleError(err)
			res.Json(w, dto.ErrorRes{
				Error: mes,
			}, code)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
