package app

import (
	"github.com/JulyInSummer/cinematic/internal/app/service"
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres"
	"github.com/JulyInSummer/cinematic/internal/app/transport/http"
	v1 "github.com/JulyInSummer/cinematic/internal/app/transport/http/handlers/v1"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			postgres.NewModule(),
			service.NewModule(),
			http.NewModule(),
			v1.NewModule(),
		),
		fx.Provide(
			zap.NewProduction,
			NewConfig,
		),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	)
}
