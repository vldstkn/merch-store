package account

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
	lis, err := net.Listen("tcp", app.Config.Addresses.Account)
	if err != nil {
		app.Logger.Error(err.Error(),
			slog.String("Error location", "net.Listen"),
			slog.String("Account address", app.Config.Addresses.Account),
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
		app.Logger.Error(err.Error(),
			slog.String("Error location", "NewHandler"),
		)
		return err
	}
	server := grpc.NewServer(opts...)
	defer server.Stop()
	pb.RegisterAccountServer(server, handler)
	app.Logger.Info("Server start",
		slog.String("Name", "Account"),
		slog.String("Address", app.Config.Addresses.Account),
		slog.String("Mode", app.Mode),
	)
	err = server.Serve(lis)
	if err != nil {
		app.Logger.Error(err.Error(),
			slog.String("Error location", "server.Serve"),
			slog.String("Account address", app.Config.Addresses.Account),
		)
		return err
	}
	return nil
}
