package db

import (
	"fmt"
	"log"

	"rest_go/configs"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDataBase(env string) {
	var err error
	var envFile string

	if env == "test" {
		envFile = "../../.env.test"
	} else {
		envFile = ".env." + env
	}
	configs.LoadEnv(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DBURL := configs.GetDatabaseURL()

	DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database")
	}

	logMode := configs.GetEnv("DB_LOG_MODE")
	DB.LogMode(logMode == "true")
}
