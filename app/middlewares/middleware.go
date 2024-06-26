package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !skipAuth[c.FullPath()] {
			AuthorizeRequest(c)
		}
		c.Next()
	}
}

var skipAuth = map[string]bool{
	"/api/register": true,
	"/api/login":    true,
	"/swagger/*any": true,
}
