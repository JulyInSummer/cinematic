package middlewares

import (
	"github.com/JulyInSummer/cinematic/internal/app/pkg/jwt"
	"github.com/JulyInSummer/cinematic/internal/app/pkg/rest"
	"github.com/JulyInSummer/cinematic/internal/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(conf *service.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, rest.Unauthorized())
			c.Abort()
			return
		}

		token := tokenString[len("Bearer "):]

		err := jwt.VerifyToken(token, conf.Secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, rest.InvalidToken())
			c.Abort()
			return
		}

		c.Next()
	}
}
