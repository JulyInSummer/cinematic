package v1

import (
	"errors"
	"github.com/JulyInSummer/cinematic/internal/app/pkg/rest"
	"github.com/JulyInSummer/cinematic/internal/app/service"
	"github.com/JulyInSummer/cinematic/internal/app/transport/http/handlers/v1/request"
	"github.com/JulyInSummer/cinematic/internal/app/transport/http/handlers/v1/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
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

// @Security ApiKeyAuth
// @Router /movies/create [post]
// @Summary Create A Single Movie
// @Description Create A Single Movie
// @Tags movies
// @Accept json
// @Produce json
// @Success 200 {object} rest.APIResponse
// @Failure 422 {object} rest.APIError
// @Failure 500 {object} rest.APIError
func (h *HandlerV1) CreateMovie(c *gin.Context) error {
	var req request.CreateMovie
	method := "HandlerV1.CreateMovie"

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(method, zap.Error(err))
		return rest.InvalidRequestData(err)
	}

	h.logger.Info(method, zap.Any("request", req))

	if err := h.service.Create(c.Request.Context(), req.ToDomain()); err != nil {
		h.logger.Error(method, zap.Error(err))
		return err
	}

	c.JSON(http.StatusCreated, gin.H{"message": "movie created successfully"})
	return nil
}

// @Security ApiKeyAuth
// @Router /movies [get]
// @Summary Get All Movies
// @Description Create All Movies
// @Tags movies
// @Accept json
// @Produce json
// @Success 200 {object} []response.Movie
// @Failure 422 {object} rest.APIError
// @Failure 500 {object} rest.APIError
func (h *HandlerV1) GetAll(c *gin.Context) error {
	method := "HandlerV1.GetAll"

	h.logger.Info(method, zap.Any("request", c.Request.URL.RawQuery))

	movies, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		h.logger.Error(method, zap.Error(err))
		return err
	}

	var resp []response.Movie
	for _, movie := range movies {
		resp = append(resp, response.Movie{
			ID:       movie.ID,
			Title:    movie.Title,
			Director: movie.Director,
			Year:     movie.Year,
			Plot:     movie.Plot,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
	return nil
}

// @Security ApiKeyAuth
// @Router /movies/{id} [get]
// @Summary Get Movie By ID
// @Description Get Movie By ID
// @Tags movies
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.Movie
// @Failure 422 {object} rest.APIError
// @Failure 500 {object} rest.APIError
func (h *HandlerV1) GetByID(c *gin.Context) error {
	method := "HandlerV1.GetByID"

	h.logger.Info(method, zap.Any("request", c.Request.URL.RawQuery))

	movieID, err := rest.ParseIntParam(c, "id")
	if err != nil {
		h.logger.Error(method, zap.Error(err))
		return rest.InvalidRequestData(err)
	}

	movie, err := h.service.GetByID(c.Request.Context(), movieID)
	if err != nil {
		h.logger.Error(method, zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rest.NotFound()
		}
		return err
	}

	resp := response.Movie{
		ID:       movie.ID,
		Title:    movie.Title,
		Director: movie.Director,
		Year:     movie.Year,
		Plot:     movie.Plot,
	}

	c.JSON(http.StatusOK, resp)
	return nil
}

// @Security ApiKeyAuth
// @Router /movies [put]
// @Summary Update Movie
// @Description Update Movie
// @Tags movies
// @Accept json
// @Produce json
// @Param entity body request.UpdateMovie true "entity"
// @Success 200 {object} rest.APIResponse
// @Failure 422 {object} rest.APIError
// @Failure 500 {object} rest.APIError
func (h *HandlerV1) Update(c *gin.Context) error {
	method := "HandlerV1.UpdateMovie"
	var req request.UpdateMovie

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(method, zap.Error(err))
		return rest.InvalidJSON()
	}

	h.logger.Info(method, zap.Any("request", req))

	err := h.service.Update(c.Request.Context(), req.ToDomain())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rest.NotFound()
		}
		h.logger.Error(method, zap.Error(err))
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "movie updated successfully"})
	return nil
}

// @Security ApiKeyAuth
// @Router /movies/{id} [delete]
// @Summary Delete Movie
// @Description Delete Movie
// @Tags movies
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} rest.APIResponse
// @Failure 422 {object} rest.APIError
// @Failure 500 {object} rest.APIError
func (h *HandlerV1) Delete(c *gin.Context) error {
	method := "HandlerV1.Delete"

	movieID, err := rest.ParseIntParam(c, "id")
	if err != nil {
		h.logger.Error(method, zap.Error(err))
		return rest.InvalidRequestData(err)
	}

	if err = h.service.Delete(c.Request.Context(), movieID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rest.NotFound()
		}
		h.logger.Error(method, zap.Error(err))
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "movie deleted successfully"})
	return nil
}
