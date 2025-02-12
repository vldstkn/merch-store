package account

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

func (repo *Repository) GetByName(name string) *models.User {
	var user models.User
	err := repo.DB.Get(&user, `SELECT * FROM users WHERE name=$1`, name)
	if err != nil {
		return nil
	}
	return &user
}
func (repo *Repository) GetById(id int64) *models.User {
	var user models.User
	err := repo.DB.Get(&user, `SELECT * FROM users WHERE id=$1`, id)
	if err != nil {
		return nil
	}
	return &user
}
func (repo *Repository) Create(user *models.User) (int64, error) {
	var id int64
	err := repo.DB.QueryRow(`INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id`,
		user.Name, user.Password).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (repo *Repository) DeductBalance(userId, amount int64) error {
	_, err := repo.DB.Exec(`UPDATE users SET balance=balance-$2 WHERE id=$1`, userId, amount)
	return err
}
func (repo *Repository) Refund(userId, amount int64) error {
	_, err := repo.DB.Exec(`UPDATE users SET balance=balance+$2 WHERE id=$1`, userId, amount)
	return err
}
