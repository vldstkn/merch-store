package products

import (
	"google.golang.org/grpc"
	"log/slog"
	"merch_store/internal/config"
	"merch_store/pkg/db"
	"merch_store/pkg/pb"
	"net"
)

type AppDeps struct {
	Config *config.Config
	DB     *db.DB
	Logger *slog.Logger
	Mode   string
}

type App struct {
	Config *config.Config
	DB     *db.DB
	Logger *slog.Logger
	Mode   string
}

func NewApp(deps *AppDeps) *App {
	return &App{
		Config: deps.Config,
		DB:     deps.DB,
		Logger: deps.Logger,
		Mode:   deps.Mode,
	}
}

func (app *App) Run() error {
	var opts []grpc.ServerOption
	lis, err := net.Listen("tcp", app.Config.Addresses.Products)
	if err != nil {
		app.Logger.Error(err.Error(),
			slog.String("Error location", "net.Listen"),
			slog.String("Products address", app.Config.Addresses.Products),
		)
		return err
	}
	defer lis.Close()
	repository := NewRepository(app.DB)
	service := NewService(&ServiceDeps{
		Logger:     app.Logger,
		Repository: repository,
	})
	handler, err := NewHandler(&HandlerDeps{
		Config:  app.Config,
		Logger:  app.Logger,
		Service: service,
	})
	if err != nil {
		return err
	}
	server := grpc.NewServer(opts...)
	pb.RegisterProductServer(server, handler)
	app.Logger.Info("Server start",
		slog.String("Name", "Products"),
		slog.String("Address", app.Config.Addresses.Products),
		slog.String("Mode", app.Mode),
	)
	err = server.Serve(lis)
	defer server.Stop()
	if err != nil {
		app.Logger.Error(err.Error(),
			slog.String("Error location", "server.Serve"),
			slog.String("Products address", app.Config.Addresses.Products),
			slog.String("Mode", app.Mode),
		)
		return err
	}
	return nil
}
