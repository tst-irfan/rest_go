package tests

import (
	"log"
	"rest_go/app/models"
)

func Setup() {
	env := "test"
	models.ConnectDataBase(env)
}

func Teardown() {
	models.DB.Exec("TRUNCATE users CASCADE")
	models.DB.Exec("TRUNCATE profiles CASCADE")
	models.DB.Close()
	log.Print("Teardown complete")
}
