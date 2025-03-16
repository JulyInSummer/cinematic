package v1

import (
	"github.com/JulyInSummer/cinematic/internal/app/pkg/rest"
	"github.com/JulyInSummer/cinematic/internal/app/transport/http/handlers/v1/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// @Router /register [post]
// @Summary Create A New User
// @Description Create A New User
// @Tags auth
// @Accept json
// @Produce json
// @Param entity body request.CreateUser true "entity"
// @Success 200 {object} rest.APIResponse
// @Failure 422 {object} rest.APIError
// @Failure 500 {object} rest.APIError
func (h *HandlerV1) CreateUser(c *gin.Context) error {
	method := "HandlerV1.CreateUser"

	var data request.CreateUser
	if err := c.ShouldBindJSON(&data); err != nil {
		h.logger.Error(method, zap.Error(err))
		return rest.InvalidRequestData(err)
	}

	err := h.service.CreateUser(c.Request.Context(), data.ToDomain())
	if err != nil {
		h.logger.Error(method, zap.Error(err))
		return err
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	return nil
}
