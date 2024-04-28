package factories

import (
	"rest_go/app/models"

	"github.com/brianvoe/gofakeit/v7"
)

func GenerateUser() *models.User {
	userParams := models.User{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, false, 14),
	}
	user, err := models.UserQuery.Create(userParams)
	if err != nil {
		panic(err)
	}
	return user
}

func GenerateUsers(quantity int) []*models.User {
	users := make([]*models.User, quantity)
	for i := 0; i < quantity; i++ {
		users[i] = GenerateUser()
	}
	return users
}
