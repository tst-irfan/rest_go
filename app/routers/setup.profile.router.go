package routers

import (
	"rest_go/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProfileRoutes(r *gin.RouterGroup) {
	r.GET("/profiles", controllers.ShowAllProfiles)
	r.GET("/profiles/:id", controllers.GetProfile)
	r.GET("/my-profile", controllers.ShowMyProfile)
	r.POST("/profiles", controllers.SaveProfile)
	r.PUT("/profiles", controllers.UpdateProfile)
	r.DELETE("/profiles/:id", controllers.DeleteProfile)
}
