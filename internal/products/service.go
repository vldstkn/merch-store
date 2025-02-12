package products

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"merch-store/internal/interfaces"
	"merch-store/internal/models"
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

func (service *Service) AddProductToInventory(userId, productId int64) error {
	err := service.Repository.AddProductToInventory(userId, productId)
	if err != nil {
		service.Logger.Error(err.Error(),
			slog.String("Error location", "Repository.AddProductToInventory"),
			slog.Int64("User Id", userId),
			slog.Int64("Product Id", productId),
		)
		return status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	return nil
}
func (service *Service) GetUserInventory(userId int64) []models.Inventory {
	inventory := service.Repository.GetUserInventory(userId)
	return inventory
}
func (service *Service) GetPriceProduct(productId int64) (int64, error) {
	product := service.Repository.GetProduct(productId)
	if product == nil {
		service.Logger.Error("product id is bad",
			slog.String("Error location", "Repository.GetProduct"),
			slog.Int64("Product Id", productId),
		)
		return 0, status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	return product.Price, nil
}
