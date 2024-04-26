package middlewares

import (
	"net/http"
	"rest_go/app/helpers"
	"rest_go/app/utils/token"

	"github.com/gin-gonic/gin"
)

func AuthorizeRequest(c *gin.Context) {
	err := token.TokenValid(c)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusUnauthorized)
		c.Abort()
		return
	}
	id, err := token.ExtractTokenID(c)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusUnauthorized)
		c.Abort()
		return
	}
	c.Set("UserID", id)
}
