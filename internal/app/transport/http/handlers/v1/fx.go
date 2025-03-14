package v1

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"api_v1",
		fx.Provide(NewHandlerV1),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.Named("api_v1")
		}),
	)
}
