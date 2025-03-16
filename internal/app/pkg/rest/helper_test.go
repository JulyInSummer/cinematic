package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseIntParamSuccess(t *testing.T) {
	input := "12343"
	ctx := &gin.Context{}
	param := []gin.Param{
		{
			Key:   "id",
			Value: input,
		},
	}

	params := gin.Params{}
	params = append(params, param...)
	ctx.Params = params
	result, err := ParseIntParam(ctx, "id")
	require.NoError(t, err)
	require.Equal(t, 12343, result)
}

func TestParseIntParamFail(t *testing.T) {
	input := "12343as"
	ctx := &gin.Context{}
	param := []gin.Param{
		{
			Key:   "id",
			Value: input,
		},
	}

	params := gin.Params{}
	params = append(params, param...)
	ctx.Params = params
	_, err := ParseIntParam(ctx, "id")
	require.Error(t, err)
}
