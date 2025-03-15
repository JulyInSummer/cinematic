package http

import (
	"github.com/JulyInSummer/cinematic/internal/app/pkg/rest"
	"github.com/JulyInSummer/cinematic/internal/app/pkg/validators"
	srvc "github.com/JulyInSummer/cinematic/internal/app/service"
	v1 "github.com/JulyInSummer/cinematic/internal/app/transport/http/handlers/v1"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("date_format", validators.DateFormat)
	}

	router := engine.Group("/api")
	router.GET("/ping", ping)

	apiV1 := v1.NewHandlerV1(logger, service)
	routerV1 := router.Group("/v1")

	{
		routerV1.GET("/movies", rest.Handle(apiV1.GetAll))
		routerV1.GET("/movies/:id", rest.Handle(apiV1.GetByID))
		routerV1.DELETE("/movies/:id", rest.Handle(apiV1.Delete))
		routerV1.PUT("/movies", rest.Handle(apiV1.Update))
		routerV1.POST("/movies/create", rest.Handle(apiV1.CreateMovie))
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
