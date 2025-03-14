package postgres

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"storage_postgres",
		fx.Provide(
			NewConfig,
			NewStorage,
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("storage_postgres")
		}),
	)
}
