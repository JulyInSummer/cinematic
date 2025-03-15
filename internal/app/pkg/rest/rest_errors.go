package rest

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"status_code"`
	Message    any `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func InvalidRequestData(err error) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    err.Error(),
	}
}

func NotFound() APIError {
	return NewAPIError(http.StatusNotFound, errors.New("not found"))
}

func Unauthorized() APIError {
	return NewAPIError(http.StatusUnauthorized, errors.New("unauthorized"))
}

func InvalidCredentials() APIError {
	return NewAPIError(http.StatusUnauthorized, errors.New("email or password is invalid"))
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, errors.New("invalid JSON data"))
}

type APIFunc func(c *gin.Context) error

func Handle(fn APIFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(c); err != nil {
			if apiErr, ok := err.(APIError); ok {
				c.JSON(apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"status_code": http.StatusInternalServerError,
					"message":     "internal server error",
				}
				c.JSON(http.StatusInternalServerError, errResp)
			}
		}
	}
}
