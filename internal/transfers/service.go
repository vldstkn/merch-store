package transfers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"merch_store/internal/interfaces"
	"merch_store/internal/models"
	"net/http"
)

type ServiceDeps struct {
	Repository interfaces.TransfersRepository
	Logger     *slog.Logger
}
type Service struct {
	Repository interfaces.TransfersRepository
	Logger     *slog.Logger
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{
		Repository: deps.Repository,
		Logger:     deps.Logger,
	}
}
func (service *Service) Create(userFrom, userTo string, amount int64) error {
	err := service.Repository.Create(userFrom, userTo, amount)
	if err != nil {
		service.Logger.Error(err.Error(),
			slog.String("Error location", "Repository.Create"),
			slog.String("User from", userFrom),
			slog.String("User to", userTo),
			slog.Int64("Amount", amount),
		)
		return status.Errorf(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	return nil
}

func (service *Service) GetHistory(userName string) models.History {
	rec := service.Repository.GetReceived(userName)
	sent := service.Repository.GetSent(userName)
	return models.History{
		Received: rec,
		Sent:     sent,
	}

}
