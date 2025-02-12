package products

import (
	"merch-store/internal/models"
	"merch-store/pkg/db"
)

type Repository struct {
	DB *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
func (repo *Repository) AddProductToInventory(userId int64, productId int64) error {
	_, err := repo.DB.Exec(`INSERT INTO products_users 
    														 (user_id, product_id) VALUES ($1, $2) 
                                 ON CONFLICT (user_id, product_id) 
                                 DO UPDATE SET quantity=products_users.quantity+1`, userId, productId)
	return err
}
func (repo *Repository) GetUserInventory(userId int64) []models.Inventory {
	var inventory []models.Inventory
	err := repo.DB.Select(&inventory, `SELECT type, quantity 
																						FROM products_users pu
																						JOIN products p on p.id = pu.product_id
																						WHERE pu.user_id=$1
																		 `, userId)
	if err != nil {
		return nil
	}
	return inventory
}
func (repo *Repository) GetProduct(productId int64) *models.Product {
	var product models.Product
	err := repo.DB.Get(&product, `SELECT price FROM products WHERE id=$1`, productId)
	if err != nil {
		return nil
	}
	return &product
}
