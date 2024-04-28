package tests

import (
	"rest_go/db"
	"rest_go/initializer"
)

func SetupTest() {
	db.ConnectDataBase("test")
	initializer.AutoMigrate()
}

func TeardownTest() {
	db.DB.Exec("TRUNCATE users RESTART IDENTITY CASCADE")
	db.DB.Exec("TRUNCATE profiles RESTART IDENTITY CASCADE")
	db.DB.Close()
}
