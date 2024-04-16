// code with the error
// fix the error by adding a semicolon at the end of the line

package main

import (
	"auth_go/app/controllers"
	"auth_go/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!!",
		})
	})
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	r.Run(":8080")
}
