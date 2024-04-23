package main

import (
	"auth_go/app/models"
	"auth_go/config"
)

func main() {

	models.ConnectDataBase()

	r := config.Router()

	r.Run(":8080")
}
