package interfaces

import "merch_store/internal/models"

type ProductsService interface {
	AddProductToInventory(userName string, productType string) error
	GetUserInventory(userName string) []models.Inventory
	GetPriceProduct(productType string) (int64, error)
}
type ProductsRepository interface {
	AddProductToInventory(userName string, productType string) error
	GetUserInventory(userName string) []models.Inventory
	GetProduct(productType string) *models.Product
}
