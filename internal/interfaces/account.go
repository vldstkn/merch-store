package interfaces

import (
	"merch_store/internal/models"
	"merch_store/pkg/jwt"
)

type AccountService interface {
	RegisterOrLogin(userName, password string) (int64, error)
	IssueToken(secret string, data jwt.Data) (string, error)
	UserIsExists(userName string) bool
	DeductBalance(userName string, amount int64) error
	Refund(userName string, amount int64) error
	GetBalanceById(userName string) (int64, error)
	TransferCoins(userFromName, userToName string, amount int64) error
}
type AccountRepository interface {
	GetByName(name string) *models.User
	Create(user *models.User) (int64, error)
	DeductBalance(userName string, amount int64) error
	Refund(userName string, amount int64) error
	TransferCoins(userFromName, userToName string, amount int64) error
}
