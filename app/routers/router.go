package routers

import (
	"rest_go/app/middlewares"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Middleware())
	public := r.Group("/api")

	SetupAuthRoutes(public)
	SetupProfileRoutes(public)

	return r
}
