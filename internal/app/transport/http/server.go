package http

import (
	srvc "github.com/JulyInSummer/cinematic/internal/app/service"
	v1 "github.com/JulyInSummer/cinematic/internal/app/transport/http/handlers/v1"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	server  *http.Server
	service srvc.ServiceI
	logger  *zap.Logger
	config  *Config
}

func NewHTTPServer(config *Config, logger *zap.Logger, service srvc.ServiceI) *Server {
	if config.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	router := engine.Group("/api")
	router.GET("/ping", ping)

	apiV1 := v1.NewHandlerV1(logger, service)
	routerV1 := router.Group("/v1")

	{
		routerV1.GET("/ping", apiV1.Ping)
	}

	return &Server{
		server: &http.Server{
			Addr:        config.Port,
			Handler:     engine,
			ReadTimeout: config.ReadTimeout * time.Second,
		},
		logger: logger,
		config: config,
	}
}

func ping(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
