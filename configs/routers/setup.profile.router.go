package routers

import (
	"auth_go/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProfileRoutes(r *gin.RouterGroup) {
	r.GET("/profiles", controllers.ShowAllProfiles)
	r.GET("/profiles/:id", controllers.GetProfile)
	r.POST("/profiles", controllers.SaveProfile)
	r.PUT("/profiles/:id", controllers.UpdateProfile)
	r.DELETE("/profiles/:id", controllers.DeleteProfile)
}
