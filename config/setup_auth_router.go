package config

import (
	"auth_go/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}
