package v1

import (
	"github.com/JulyInSummer/cinematic/internal/app/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HandlerV1 struct {
	logger  *zap.Logger
	service service.ServiceI
}

func NewHandlerV1(logger *zap.Logger, service service.ServiceI) *HandlerV1 {
	return &HandlerV1{
		logger:  logger,
		service: service,
	}
}

func (h *HandlerV1) Ping(c *gin.Context) {
	method := "handlerV1.ping"
	h.logger.Info(method, zap.String("url", c.Request.URL.Path))
	c.JSON(200, gin.H{"message": "pong"})
}
