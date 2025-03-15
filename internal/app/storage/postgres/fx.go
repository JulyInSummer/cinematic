package postgres

import (
	"fmt"
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/models"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewModule() fx.Option {
	return fx.Module(
		"storage_postgres",
		fx.Provide(
			func(conf *Config) (*gorm.DB, error) {
				db, err := gorm.Open(
					postgres.New(postgres.Config{
						DriverName: "postgres",
						DSN:        getConnString(conf),
					}),
				)
				if err != nil {
					return nil, fmt.Errorf("failed to connect to postgres database: %v", err)
				}

				err = db.AutoMigrate(&models.Movie{}, &models.User{})
				if err != nil {
					return nil, fmt.Errorf("failed to auto migrate: %w", err)
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
