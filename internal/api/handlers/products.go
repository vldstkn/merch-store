package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"merch_store/internal/api/dto"
	"merch_store/internal/api/middleware"
	"merch_store/internal/config"
	http_errors "merch_store/internal/errors"
	"merch_store/internal/models"
	grpc_conn "merch_store/pkg/grpc-conn"
	"merch_store/pkg/pb"
	"merch_store/pkg/res"
	"net/http"
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

func NewProductsHandler(router chi.Router, deps *ProductsHandlerDeps) error {
	productsConn, err := grpc_conn.NewClientConn(deps.Config.Addresses.Products)
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
		r.Use(middleware.IsAuthed(handler.Config.Auth.Jwt))
		r.Get("/buy/{item}", handler.Buy())
	})
	return nil
}

// Buy
//
//		@Summary		Купить предмет за монеты.
//		@ID				buy
//		@Produce		json
//		@Param		productType	path string	true	"Тип продукта, который покупается."
//		@Success		200 		"Успешный ответ"
//		@Failure		400		{object}	dto.ErrorRes	"Неверный запрос."
//		@Failure		401		{object}	dto.ErrorRes	"Пользователь не авторизован."
//		@Failure		500		{object}	dto.ErrorRes	"Внутренняя ошибка сервера."
//		@Router			/buy/{productType} [get]
//	 @Security BearerAuth
func (handler *ProductsHandler) Buy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productType := chi.URLParam(r, "item")
		isValid := models.IsValid(productType)
		if !isValid {
			handler.Logger.Error("product type undefined",
				slog.String("Error location", "chi.URLParam"),
				slog.String("Item type", productType),
			)
			res.Json(w, dto.ErrorRes{
				Error: http.StatusText(http.StatusBadRequest),
			}, http.StatusBadRequest)
			return
		}
		name := r.Context().Value("authData").(middleware.AuthData).Name
		_, err := handler.ProductsClient.Buy(context.Background(), &pb.BuyReq{
			UserName:    name,
			ProductType: productType,
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
