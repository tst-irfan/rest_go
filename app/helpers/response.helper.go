package helpers

import (
	"auth_go/app/types"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, types.Error{
		Success: false,
		Message: message,
	})
}

func ResponseSuccess[T any](c *gin.Context, message string, data T, statusCode int) {
	c.JSON(statusCode, types.Success[T]{
		Success: true,
		Message: message,
		Data:    data,
	})
}
