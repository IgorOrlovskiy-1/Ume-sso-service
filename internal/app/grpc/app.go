package grpc

import (
	"net"
	"fmt"
	"log/slog"
	"google.golang.org/grpc"
	authgrpc "github.com/IgorOrlovskiy-1/Ume-sso-service/internal/grpc/auth"
)

type App struct {
	log *slog.Logger
	gRPCServer *grpc.Server
	port int
}

func New( log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer)

	return &App{
	log: log,
	gRPCServer: gRPCServer,
	port: port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpc.app.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	log.Info("Starting gRPC server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpc.app.Stop"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	log.Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}