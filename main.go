package main

import (
	"auth_go/app/models"
	. "auth_go/configs"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	env := os.Getenv("ENV")
	models.ConnectDataBase(env)

	r := Router()

	r.Run(":8080")
}
