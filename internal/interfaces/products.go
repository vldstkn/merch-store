package interfaces

import "merch-store/internal/models"

type ProductsService interface {
	AddProductToInventory(userId, productId int64) error
	GetUserInventory(userId int64) []models.Inventory
	GetPriceProduct(productId int64) (int64, error)
}
type ProductsRepository interface {
	AddProductToInventory(userId int64, productId int64) error
	GetUserInventory(userId int64) []models.Inventory
	GetProduct(productId int64) *models.Product
}
