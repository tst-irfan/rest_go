package main

import (
	"os"
	"rest_go/app/models"
	. "rest_go/configs"

	"github.com/joho/godotenv"
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
