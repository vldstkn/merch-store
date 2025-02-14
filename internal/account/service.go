package account

import (
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	http_errors "merch_store/internal/errors"
	"merch_store/internal/interfaces"
	"merch_store/internal/models"
	"merch_store/pkg/jwt"
	"net/http"
	"time"
)

type ServiceDeps struct {
	Repository interfaces.AccountRepository
	Logger     *slog.Logger
}
type Service struct {
	Repository interfaces.AccountRepository
	Logger     *slog.Logger
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{
		Repository: deps.Repository,
		Logger:     deps.Logger,
	}
}

func (service *Service) RegisterOrLogin(userName, password string) (int64, error) {
	user := service.Repository.GetByName(userName)
	if user != nil {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return -1, status.Errorf(codes.InvalidArgument, http_errors.InvalidNameOrPassword)
		}
		return user.Id, nil
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, status.Errorf(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	user = &models.User{
		Name:     userName,
		Password: string(hashPassword),
	}
	id, err := service.Repository.Create(user)
	if err != nil {
		service.Logger.Error(err.Error(),
			slog.String("Error location", "Repository.Create"),
			slog.String("User name", userName),
		)
		return -1, status.Errorf(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	return id, nil
}
func (service *Service) IssueToken(secret string, data jwt.Data) (string, error) {
	j := jwt.NewJWT(secret)
	token, err := j.Create(data, time.Now().Add(time.Hour*2).Add(time.Minute*10))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service *Service) UserIsExists(userName string) bool {
	user := service.Repository.GetByName(userName)
	return user != nil
}

func (service *Service) DeductBalance(userName string, amount int64) error {
	user := service.Repository.GetByName(userName)
	if user == nil {
		return status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	err := service.Repository.DeductBalance(userName, amount)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, http_errors.InsufficientFunds)
	}
	return err
}

func (service *Service) Refund(userName string, amount int64) error {
	user := service.Repository.GetByName(userName)
	if user == nil {
		return status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	err := service.Repository.Refund(userName, amount)
	if err != nil {
		return status.Errorf(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	return err

}

func (service *Service) GetBalanceById(userName string) (int64, error) {
	user := service.Repository.GetByName(userName)
	if user == nil {
		return 0, status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	return user.Balance, nil
}
func (service *Service) TransferCoins(userFromName, userToName string, amount int64) error {
	userFrom := service.Repository.GetByName(userFromName)
	if userFrom == nil {
		service.Logger.Error(http_errors.UserNameNotFound,
			slog.String("User name from", userFromName),
		)
		return status.Errorf(codes.InvalidArgument, http_errors.UserNameNotFound)
	}
	if userFrom.Balance < amount {
		service.Logger.Error(http_errors.InsufficientFunds,
			slog.String("User name from", userFromName),
			slog.Int64("Amount", amount),
		)
		return status.Errorf(codes.InvalidArgument, http_errors.InsufficientFunds)
	}
	userTo := service.Repository.GetByName(userToName)
	if userTo == nil {
		service.Logger.Error(http_errors.UserNameNotFound,
			slog.String("User name to", userToName),
		)
		return status.Errorf(codes.InvalidArgument, http_errors.UserNameNotFound)
	}
	err := service.Repository.TransferCoins(userFromName, userToName, amount)
	if err != nil {
		service.Logger.Error(err.Error(),
			slog.String("User name from", userFromName),
			slog.String("User name to", userToName),
			slog.Int64("User name amount", amount),
		)
		return status.Errorf(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	return nil
}
