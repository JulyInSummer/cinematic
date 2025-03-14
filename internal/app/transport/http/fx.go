package http

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
)

func NewModule() fx.Option {
	return fx.Module(
		"http_server",
		fx.Provide(
			NewConfig,
			NewHTTPServer,
		),
		fx.Invoke(func(lc fx.Lifecycle, server *Server) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					ln, err := net.Listen("tcp", server.server.Addr)
					if err != nil {
						return fmt.Errorf("failed to listen: %w", err)
					}

					go server.server.Serve(ln)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return server.server.Shutdown(ctx)
				},
			})
		}),
		fx.Decorate(func(logger *zap.Logger) *zap.Logger {
			return logger.Named("http_server")
		}),
	)
}
