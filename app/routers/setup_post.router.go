package routers

import (
	"rest_go/app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.RouterGroup) {
	controller := controllers.NewPostController()
	r.GET("/posts", controller.GetAllPosts)
	r.GET("/posts/:id", controller.GetPost)
	r.POST("/posts", controller.CreatePost)
	r.PUT("/posts/:id", controller.UpdatePost)
	r.DELETE("/posts/:id", controller.DeletePost)
}

