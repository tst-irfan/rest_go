package models_test

import (
	"rest_go/app/models"
	"rest_go/tests"
	"rest_go/tests/factories"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/bcrypt"
)

func TestSaveUser(t *testing.T) {
	tests.Setup()

	password := gofakeit.Password(true, true, true, false, false, 12)
	userParams := models.User{
		Email:    gofakeit.Email(),
		Password: password,
	}

	user, err := userParams.SaveUser()
	if err != nil {
		t.Errorf("Error occured while saving user: %v", err)
	}
	if user.ID == 0 {
		t.Errorf("Expected user ID to be greater than 0, but got %v", user.ID)
	}
	if user.Email != userParams.Email {
		t.Errorf("Expected email to be %v, but got %v", userParams.Email, user.Email)
	}
	if user.Password == password {
		t.Errorf("Expected password to be encrypted")
	}

	user, err = userParams.SaveUser()
	if err == nil {
		t.Errorf("Expected error to be thrown")
	}

	userParams.Email = gofakeit.Email()
	userParams.Password = ""
	user, err = userParams.SaveUser()
	if err == nil {
		t.Errorf("Expected error to be thrown")
	}

	userParams.Email = ""
	userParams.Password = password
	user, err = userParams.SaveUser()
	if err == nil {
		t.Errorf("Expected error to be thrown")
	}
	tests.Teardown()
}

func TestGetUserByEmail(t *testing.T) {
	tests.Setup()

	user := factories.GenerateUser()

	data, err := models.GetUserByEmail(user.Email)

	if err != nil {
		t.Errorf("Error occured while fetching user: %v", err)
	}

	if data.ID != user.ID {
		t.Errorf("Expected user ID to be %v, but got %v", user.ID, data.ID)
	}
	if data.Email != user.Email {
		t.Errorf("Expected email to be %v, but got %v", user.Email, data.Email)
	}

	tests.Teardown()
}

func TestVerifyPassword(t *testing.T) {
	tests.Setup()

	password := gofakeit.Password(true, true, true, false, false, 12)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Errorf("Error occured while hashing password: %v", err)
	}
	data, err := models.VerifyPassword(password, string(hashedPassword))
	if err != nil {
		t.Errorf("Error occured while verifying password: %v", err)
	}

	if data != true {
		t.Errorf("Expected password to be verified")
	}
	data2, err := models.VerifyPassword("test", string(hashedPassword))
	if data2 != false {
		t.Errorf("Expected password to be invalid")
	}
}
