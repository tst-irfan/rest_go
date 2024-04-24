package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbName, DbPassword)

	DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database")
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Profile{})
}
