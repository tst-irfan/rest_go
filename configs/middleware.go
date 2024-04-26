package configs

import (
	"auth_go/configs/middlewares"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !skipAuth[c.FullPath()] {
			middlewares.AuthorizeRequest(c)
		}
		c.Next()
	}
}

var skipAuth = map[string]bool{
	"/api/register": true,
	"/api/login":    true,
}
