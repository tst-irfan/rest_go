package initializer

import (
	"rest_go/app/routers"
	"rest_go/configs"
	"rest_go/db"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"rest_go/docs"
)

func RunApplication() {
	configs.LoadEnv(".env")
	env := configs.GetEnv("ENV")

	db.ConnectDataBase(env)
	AutoMigrate()

	docs.SwaggerInfo.Title = "Rest Go API"
	docs.SwaggerInfo.Description = "This is a sample server for a RESTful API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r := routers.Router()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
