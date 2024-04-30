package initializer

import (
	"rest_go/app/models"
	"rest_go/db"
)

func AutoMigrate() {
	db.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
	)
}
