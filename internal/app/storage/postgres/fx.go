package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"storage_postgres",
		fx.Provide(
			func(conf *Config, ctx context.Context) (*pgx.Conn, error) {
				db, err := pgx.Connect(ctx, getConnString(conf))
				if err != nil {
					return nil, fmt.Errorf("failed to connect to postgres database: %v", err)
				}

				if err = db.Ping(ctx); err != nil {
					return nil, fmt.Errorf("failed to ping postgres database: %v", err)
				}

				return db, nil
			},
			NewConfig,
			NewStorage,
			NewUserStorage,
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("storage_postgres")
		}),
	)
}
