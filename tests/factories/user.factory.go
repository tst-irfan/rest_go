package factories

import (
	"rest_go/app/models"

	"github.com/brianvoe/gofakeit/v7"
)

func GenerateUser() *models.User {
	var userParams models.User
	userParams.Email = gofakeit.Email()
	userParams.Password = gofakeit.Password(true, true, true, false, false, 12)
	user, err := userParams.SaveUser()
	if err != nil {
		panic(err)
	}
	return user
}

func GenerateUsers(n int) []models.User {
	users := []models.User{}
	for i := 0; i < n; i++ {
		users = append(users, *GenerateUser()) // Dereference the pointer before appending
	}
	return users
}
