package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type DatabaseProperties struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadEnv(file string) {
	err := godotenv.Load(file)
	if err != nil {
		panic(err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetDatabaseProperties() DatabaseProperties {
	return DatabaseProperties{
		Host:     GetEnv("DB_HOST"),
		Port:     GetEnv("DB_PORT"),
		User:     GetEnv("DB_USER"),
		Password: GetEnv("DB_PASSWORD"),
		Name:     GetEnv("DB_NAME"),
	}
}

func GetDatabaseURL() string {
	props := GetDatabaseProperties()
	return "host=" + props.Host + " port=" + props.Port + " user=" + props.User + " dbname=" + props.Name + " password=" + props.Password + " sslmode=disable"
}





