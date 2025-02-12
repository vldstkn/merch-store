package interfaces

import (
	"merch-store/internal/models"
	"merch-store/pkg/jwt"
)

type AccountService interface {
	RegisterOrLogin(userName, password string) (int64, error)
	IssueToken(secret string, data jwt.Data) (string, error)
	UserIsExists(userId int64) bool
	DeductBalance(userId, amount int64) error
	Refund(userId, amount int64) error
}
type AccountRepository interface {
	GetByName(name string) *models.User
	Create(user *models.User) (int64, error)
	GetById(id int64) *models.User
	DeductBalance(userId, amount int64) error
	Refund(userId, amount int64) error
}
