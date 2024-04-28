package initializer

import (
	"rest_go/app/routers"
	"rest_go/configs"
	"rest_go/db"
)

func RunApplication() {
	configs.LoadEnv(".env")
	env := configs.GetEnv("ENV")

	db.ConnectDataBase(env)
	AutoMigrate()

	r := routers.Router()

	r.Run(":8080")
}
