package interfaces

import "merch_store/internal/models"

type TransfersService interface {
	Create(userFrom, userTo string, amount int64) error
	GetHistory(userName string) models.History
}
type TransfersRepository interface {
	Create(userFrom, userTo string, amount int64) error
	GetReceived(userName string) []models.Received
	GetSent(userName string) []models.Sent
}
