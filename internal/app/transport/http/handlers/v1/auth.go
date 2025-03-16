package v1

import (
	"github.com/JulyInSummer/cinematic/internal/app/pkg/rest"
	"github.com/JulyInSummer/cinematic/internal/app/transport/http/handlers/v1/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// @Router /login [post]
// @Summary Get Bearer Token
// @Description Get Bearer Token
// @Tags auth
// @Accept json
// @Produce json
// @Param entity body request.LoginCredentials true "entity"
// @Success 200 {object} rest.APIResponse
// @Failure 422 {object} rest.APIError
// @Failure 500 {object} rest.APIError
func (h *HandlerV1) Login(c *gin.Context) error {
	method := "HandlerV1.Login"

	var req request.LoginCredentials
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(method, zap.Error(err))
		return rest.InvalidRequestData(err)
	}

	token, err := h.service.GetUserTokenByEmail(c.Request.Context(), req.ToDomain())
	if err != nil {
		h.logger.Error(method, zap.Error(err))
		return rest.InvalidCredentials()
	}

	c.JSON(http.StatusOK, gin.H{"token": token.Token})
	return nil
}
