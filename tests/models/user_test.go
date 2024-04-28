package models_test

import (
	"rest_go/app/models"
	"testing"
	"github.com/brianvoe/gofakeit/v7"
	"rest_go/tests"
	
)

func TestUserSave(t *testing.T) {
	tests.SetupTest()
	userParams := models.User{
		Email: 	gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, false, 14),
	}
	user, err := models.UserQuery.Create(userParams)
	if err != nil {
		t.Errorf("Error while saving user: %v", err)
	}
	if user.ID == 0 {
		t.Errorf("Error while saving user: ID is zero")
	}
	if user.Email != userParams.Email {
		t.Errorf("Error while saving user: Email is different")
	}
	if user.Password == userParams.Password {
		t.Errorf("Error while saving user: Password is not hashed")
	}
	tests.TeardownTest()
}