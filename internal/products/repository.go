package products

import (
	"merch_store/internal/models"
	"merch_store/pkg/db"
)

type Repository struct {
	DB *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
func (repo *Repository) AddProductToInventory(userName string, productType string) error {
	_, err := repo.DB.Exec(`INSERT INTO products_users 
    														 (user_name, product_type) VALUES ($1, $2) 
                                 ON CONFLICT (user_name, product_type) 
                                 DO UPDATE SET quantity=products_users.quantity+1`, userName, productType)
	return err
}
func (repo *Repository) GetUserInventory(userName string) []models.Inventory {
	var inventory []models.Inventory
	err := repo.DB.Select(&inventory, `SELECT type, quantity 
																						FROM products_users pu
																						JOIN products p on p.type = pu.product_type
																						WHERE pu.user_name=$1
																		 `, userName)
	if err != nil {
		return nil
	}
	return inventory
}
func (repo *Repository) GetProduct(productType string) *models.Product {
	var product models.Product
	err := repo.DB.Get(&product, `SELECT price FROM products WHERE type=$1`, productType)
	if err != nil {
		return nil
	}
	return &product
}
