package configs

import (
	"log"

	"auth_go/configs/middlewares"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !skipAuth[c.FullPath()] {
			middlewares.AuthorizeRequest(c)
		}
		log.Println("Before the request is processed")
		c.Next()
	}
}

var skipAuth = map[string]bool{
	"/api/register": true,
	"/api/login":    true,
}
