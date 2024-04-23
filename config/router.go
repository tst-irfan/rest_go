package config

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")

	SetupAuthRoutes(public)

	return r
}
