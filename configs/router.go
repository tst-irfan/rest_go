package configs

import (
	"rest_go/configs/routers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(Middleware())
	public := r.Group("/api")

	routers.SetupAuthRoutes(public)
	routers.SetupProfileRoutes(public)

	return r
}
