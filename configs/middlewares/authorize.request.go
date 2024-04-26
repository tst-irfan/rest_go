package middlewares

import (
	"auth_go/app/utils/token"
	"auth_go/app/helpers"
	"net/http"

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
