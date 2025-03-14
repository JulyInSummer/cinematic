package service

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"service",
		fx.Provide(NewCinematicService),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.Named("service")
		}),
	)
}
