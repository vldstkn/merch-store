package account

import (
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	http_errors "merch-store/internal/errors"
	"merch-store/internal/interfaces"
	"merch-store/internal/models"
	"merch-store/pkg/jwt"
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

func (service *Service) UserIsExists(userId int64) bool {
	user := service.Repository.GetById(userId)
	return user != nil
}

func (service *Service) DeductBalance(userId, amount int64) error {
	user := service.Repository.GetById(userId)
	if user == nil {
		return status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	err := service.Repository.DeductBalance(userId, amount)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, http_errors.InsufficientFunds)
	}
	return err
}

func (service *Service) Refund(userId, amount int64) error {
	user := service.Repository.GetById(userId)
	if user == nil {
		return status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	err := service.Repository.Refund(userId, amount)
	if err != nil {
		return status.Errorf(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	return err

}
