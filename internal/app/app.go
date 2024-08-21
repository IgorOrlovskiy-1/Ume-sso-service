package app

import (
	"log/slog"
	"github.com/IgorOrlovskiy-1/Ume-sso-service/internal/config"
	grpcapp "github.com/IgorOrlovskiy-1/Ume-sso-service/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	cfg *config.Config,
) *App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App {
		GRPCServer: grpcApp,
	}
}
