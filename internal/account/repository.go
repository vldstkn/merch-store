package account

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

func (repo *Repository) GetByName(name string) *models.User {
	var user models.User
	err := repo.DB.Get(&user, `SELECT * FROM users WHERE name=$1`, name)
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

func (repo *Repository) DeductBalance(userName string, amount int64) error {
	_, err := repo.DB.Exec(`UPDATE users SET balance=balance-$2 WHERE name=$1`, userName, amount)
	return err
}
func (repo *Repository) Refund(userName string, amount int64) error {
	_, err := repo.DB.Exec(`UPDATE users SET balance=balance+$2 WHERE name=$1`, userName, amount)
	return err
}

func (repo *Repository) TransferCoins(userFromName, userToName string, amount int64) error {
	tr, err := repo.DB.Beginx()
	if err != nil {
		return err
	}
	_, err = repo.DB.Exec(`UPDATE users SET balance=balance-$2 WHERE name=$1`, userFromName, amount)
	if err != nil {
		tr.Rollback()
		return err
	}
	_, err = repo.DB.Exec(`UPDATE users SET balance=balance+$2 WHERE name=$1`, userToName, amount)
	if err != nil {
		tr.Rollback()
		return err
	}
	err = tr.Commit()
	return err
}
