package services

import (
	"auth_go/app/models"
	"auth_go/app/types"
	"auth_go/app/utils/token"
	"errors"
)

func RegisterUser(input types.AuthRequest) (types.UserResponse, error) {
	user := models.User{}

	user.Email = input.Email
	user.Password = input.Password

	newUser, err := user.SaveUser()
	if err != nil {
		return types.UserResponse{}, err
	}

	user = *newUser

	return models.BuildUserAtributes(user), nil
}

func LoginUser(input types.AuthRequest) (string, error) {

	user, err := models.GetUserByEmail(input.Email)
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	isPasswordMatch, err := models.VerifyPassword(input.Password, user.Password)
	if err != nil {
		return "", err
	}

	if !isPasswordMatch {
		return "", errors.New("Invalid credentials")
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return token, err
	}

	return token, nil
}


