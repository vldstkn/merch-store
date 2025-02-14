package account

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"merch_store/internal/config"
	"merch_store/internal/interfaces"
	grpc_conn "merch_store/pkg/grpc-conn"
	"merch_store/pkg/jwt"
	"merch_store/pkg/pb"
	"net/http"
)

type Handler struct {
	pb.UnsafeAccountServer
	Config          *config.Config
	Logger          *slog.Logger
	Service         interfaces.AccountService
	ProductsClient  pb.ProductClient
	TransfersClient pb.TransfersClient
}

type HandlerDeps struct {
	Config  *config.Config
	Logger  *slog.Logger
	Service interfaces.AccountService
}

func NewHandler(deps *HandlerDeps) (*Handler, error) {
	productsConn, err := grpc_conn.NewClientConn(deps.Config.Addresses.Products)
	if err != nil {
		return nil, err
	}
	productsClient := pb.NewProductClient(productsConn)

	transfersConn, err := grpc_conn.NewClientConn(deps.Config.Addresses.Products)
	if err != nil {
		return nil, err
	}
	transfersClient := pb.NewTransfersClient(transfersConn)
	return &Handler{
		Config:          deps.Config,
		Logger:          deps.Logger,
		Service:         deps.Service,
		ProductsClient:  productsClient,
		TransfersClient: transfersClient,
	}, nil
}

func (handler *Handler) Auth(ctx context.Context, r *pb.AuthReq) (*pb.AuthRes, error) {
	_, err := handler.Service.RegisterOrLogin(r.Username, r.Password)
	if err != nil {
		return nil, err
	}
	token, err := handler.Service.IssueToken(handler.Config.Auth.Jwt, jwt.Data{
		Name: r.Username,
	})
	if err != nil {
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}
	return &pb.AuthRes{
		Token: token,
	}, nil
}
func (handler *Handler) GetInfo(ctx context.Context, r *pb.GetInfoReq) (*pb.GetInfoRes, error) {
	balance, err := handler.Service.GetBalanceById(r.UserName)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, http.StatusText(http.StatusBadRequest))
	}
	resProd, err := handler.ProductsClient.GetUserInventory(context.Background(), &pb.GetUserInventoryReq{
		UserName: r.UserName,
	})
	if err != nil {
		handler.Logger.Error(err.Error(),
			slog.String("Error location", "GetInfo: ProductsClient.GetUserInventory"),
			slog.String("Layer", "Handler"),
		)
		return nil, err
	}
	resTran, err := handler.TransfersClient.GetHistory(context.Background(), &pb.GetHistoryReq{
		UserName: r.UserName,
	})
	if err != nil {
		handler.Logger.Error(err.Error(),
			slog.String("Error location", "GetInfo: TransfersClient.GetHistory"),
			slog.String("Layer", "Handler"),
		)
		return nil, err
	}
	return &pb.GetInfoRes{
		Inventory:    resProd.Inventory,
		Coins:        balance,
		CoinsHistory: resTran.CoinsHistory,
	}, nil
}
func (handler *Handler) DeductBalance(ctx context.Context, r *pb.DeductBalanceReq) (*pb.DeductBalanceRes, error) {
	err := handler.Service.DeductBalance(r.UserName, r.Amount)
	if err != nil {
		return nil, err
	}
	return &pb.DeductBalanceRes{}, nil
}
func (handler *Handler) Refund(ctx context.Context, r *pb.RefundReq) (*pb.RefundRes, error) {
	err := handler.Service.Refund(r.UserName, r.Amount)
	if err != nil {
		return nil, err
	}
	return &pb.RefundRes{}, nil
}

func (handler *Handler) TransferCoins(ctx context.Context, r *pb.TransferCoinsReq) (*pb.TransferCoinsRes, error) {
	err := handler.Service.TransferCoins(r.UserFromName, r.UserToName, r.Amount)
	if err != nil {
		return nil, err
	}
	return &pb.TransferCoinsRes{}, nil
}
