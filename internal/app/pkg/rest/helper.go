package rest

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseIntParam(c *gin.Context, param string) (uint64, error) {
	p, err := strconv.Atoi(c.Param(param))
	if err != nil {
		return 0, err
	}

	return uint64(p), nil
}
