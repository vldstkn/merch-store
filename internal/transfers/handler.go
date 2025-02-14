package transfers

import (
	"context"
	"log/slog"
	"merch_store/internal/config"
	"merch_store/internal/interfaces"
	"merch_store/internal/transfers/saga"
	grpc_conn "merch_store/pkg/grpc-conn"
	"merch_store/pkg/pb"
)

type Handler struct {
	pb.UnsafeTransfersServer
	Config        *config.Config
	Logger        *slog.Logger
	Service       interfaces.TransfersService
	AccountClient pb.AccountClient
}

type HandlerDeps struct {
	Config  *config.Config
	Logger  *slog.Logger
	Service interfaces.TransfersService
}

func NewHandler(deps *HandlerDeps) (*Handler, error) {
	accountConn, err := grpc_conn.NewClientConn(deps.Config.Addresses.Account)
	if err != nil {
		deps.Logger.Error(err.Error(),
			slog.String("Error location", "grpc_conn.NewClientConn"),
			slog.String("Account address", deps.Config.Addresses.Account),
		)
		return nil, err
	}
	accountClient := pb.NewAccountClient(accountConn)
	return &Handler{
		Config:        deps.Config,
		Logger:        deps.Logger,
		Service:       deps.Service,
		AccountClient: accountClient,
	}, nil
}

func (handler *Handler) SendCoins(ctx context.Context, r *pb.SendCoinsReq) (*pb.SendCoinsRes, error) {
	saga := saga.NewSendSaga(&saga.SendSagaDeps{
		AccountClient:   handler.AccountClient,
		TransferService: handler.Service,
	})
	err := saga.Start(r.FromUser, r.ToUser, r.Amount)
	if err != nil {
		return nil, err
	}
	return &pb.SendCoinsRes{}, nil
}
func (handler *Handler) GetHistory(ctx context.Context, r *pb.GetHistoryReq) (*pb.GetHistoryRes, error) {
	history := handler.Service.GetHistory(r.UserName)
	return &pb.GetHistoryRes{
		CoinsHistory: HistoryFromModelToGrpc(history),
	}, nil
}
