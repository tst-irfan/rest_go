package main

import (
	"auth_go/app/models"
	"auth_go/configs"
)

func main() {

	models.ConnectDataBase()

	r := configs.Router()

	r.Run(":8080")
}
