package transfers

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

func (repo *Repository) Create(userFrom, userTo string, amount int64) error {
	_, err := repo.DB.Exec(`INSERT INTO transfers (from_user_name, to_user_name, amount) VALUES ($1,$2,$3)`,
		userFrom, userTo, amount)
	return err
}

func (repo *Repository) GetReceived(userName string) []models.Received {
	var rec []models.Received

	err := repo.DB.Select(&rec, `SELECT from_user_name, amount FROM transfers WHERE to_user_name=$1`, userName)
	if err != nil {
		return nil
	}
	return rec
}
func (repo *Repository) GetSent(userName string) []models.Sent {
	var sent []models.Sent
	err := repo.DB.Select(&sent, `SELECT to_user_name, amount FROM transfers WHERE from_user_name=$1`, userName)
	if err != nil {
		return nil
	}
	return sent
}
