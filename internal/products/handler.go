package products

import (
	"context"
	"log/slog"
	"merch_store/internal/config"
	"merch_store/internal/interfaces"
	"merch_store/internal/products/saga"
	grpc_conn "merch_store/pkg/grpc-conn"
	"merch_store/pkg/pb"
)

type Handler struct {
	pb.UnsafeProductServer
	Config        *config.Config
	Logger        *slog.Logger
	Service       interfaces.ProductsService
	AccountClient pb.AccountClient
}

type HandlerDeps struct {
	Config  *config.Config
	Logger  *slog.Logger
	Service interfaces.ProductsService
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

func (handler *Handler) GetUserInventory(ctx context.Context, r *pb.GetUserInventoryReq) (*pb.GetUserInventoryRes, error) {
	inventory := handler.Service.GetUserInventory(r.UserName)
	return &pb.GetUserInventoryRes{
		Inventory: InventoryFromModelsToGrpc(inventory),
	}, nil
}
func (handler *Handler) Buy(ctx context.Context, r *pb.BuyReq) (*pb.BuyRes, error) {
	saga := saga.NewBuySaga(&saga.BuySagaDeps{
		AccountClient:  handler.AccountClient,
		ProductService: handler.Service,
	})
	err := saga.Start(r.UserName, r.ProductType)
	if err != nil {
		return nil, err
	}
	return &pb.BuyRes{}, nil
}
func (handler *Handler) GetPriceProduct(ctx context.Context, r *pb.BuyReq) (*pb.BuyRes, error) {
	return nil, nil
}
