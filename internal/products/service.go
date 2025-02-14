package products

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"merch_store/internal/interfaces"
	"merch_store/internal/models"
	"net/http"
)

type ServiceDeps struct {
	Repository interfaces.ProductsRepository
	Logger     *slog.Logger
}
type Service struct {
	Repository interfaces.ProductsRepository
	Logger     *slog.Logger
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{
		Repository: deps.Repository,
		Logger:     deps.Logger,
	}
}

func (service *Service) AddProductToInventory(userName, productType string) error {
	err := service.Repository.AddProductToInventory(userName, productType)
	if err != nil {
		service.Logger.Error(err.Error(),
			slog.String("Error location", "Repository.AddProductToInventory"),
			slog.String("User name", userName),
			slog.String("Product type", productType),
		)
		return status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	return nil
}
func (service *Service) GetUserInventory(userName string) []models.Inventory {
	inventory := service.Repository.GetUserInventory(userName)
	return inventory
}
func (service *Service) GetPriceProduct(productType string) (int64, error) {
	product := service.Repository.GetProduct(productType)
	if product == nil {
		service.Logger.Error("product id is bad",
			slog.String("Error location", "Repository.GetProduct"),
			slog.String("Product type", productType),
		)
		return 0, status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	return product.Price, nil
}
