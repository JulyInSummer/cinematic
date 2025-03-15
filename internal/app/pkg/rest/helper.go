package rest

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseIntParam(c *gin.Context, param string) (int, error) {
	p, err := strconv.Atoi(c.Param(param))
	if err != nil {
		return 0, err
	}

	return p, nil
}
