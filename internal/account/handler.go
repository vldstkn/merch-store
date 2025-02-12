package account

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"merch-store/internal/config"
	"merch-store/internal/interfaces"
	grpc_conn "merch-store/pkg/grpc-conn"
	"merch-store/pkg/jwt"
	"merch-store/pkg/pb"
	"net/http"
)

type Handler struct {
	pb.UnsafeAccountServer
	Config         *config.Config
	Logger         *slog.Logger
	Service        interfaces.AccountService
	ProductsClient pb.ProductClient
}

type HandlerDeps struct {
	Config  *config.Config
	Logger  *slog.Logger
	Service interfaces.AccountService
}

func NewHandler(deps *HandlerDeps) (*Handler, error) {
	productsConn, err := grpc_conn.NewClientConn(deps.Config.ProductsAddress)
	if err != nil {
		return nil, err
	}
	productsClient := pb.NewProductClient(productsConn)
	return &Handler{
		Config:         deps.Config,
		Logger:         deps.Logger,
		Service:        deps.Service,
		ProductsClient: productsClient,
	}, nil
}

func (handler *Handler) Auth(ctx context.Context, r *pb.AuthReq) (*pb.AuthRes, error) {
	id, err := handler.Service.RegisterOrLogin(r.Username, r.Password)
	if err != nil {
		return nil, err
	}
	token, err := handler.Service.IssueToken(handler.Config.JWTSecret, jwt.Data{
		Id: id,
	})
	if err != nil {
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}
	return &pb.AuthRes{
		Token: token,
	}, nil
}
func (handler *Handler) GetInfo(ctx context.Context, r *pb.GetInfoReq) (*pb.GetInfoRes, error) {
	userIsExists := handler.Service.UserIsExists(r.UserId)
	if !userIsExists {
		return nil, status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	response, err := handler.ProductsClient.GetUserInventory(context.Background(), &pb.GetUserInventoryReq{
		UserId: r.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetInfoRes{
		Inventory: response.Inventory,
	}, nil
}
func (handler *Handler) DeductBalance(ctx context.Context, r *pb.DeductBalanceReq) (*pb.DeductBalanceRes, error) {
	err := handler.Service.DeductBalance(r.UserId, r.Amount)
	if err != nil {
		return nil, err
	}
	return &pb.DeductBalanceRes{}, nil
}
func (handler *Handler) Refund(ctx context.Context, r *pb.RefundReq) (*pb.RefundRes, error) {
	err := handler.Service.Refund(r.UserId, r.Amount)
	if err != nil {
		return nil, err
	}
	return &pb.RefundRes{}, nil
}
