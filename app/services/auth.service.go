package services

import (
	"auth_go/app/models"
	. "auth_go/app/types"
	"auth_go/app/utils/token"
	"errors"
)

func RegisterUser(input AuthRequest) (UserResponse, error) {
	user := models.User{}

	user.Email = input.Email
	user.Password = input.Password

	newUser, err := user.SaveUser()
	if err != nil {
		return UserResponse{}, err
	}

	user = *newUser

	return buildUserAtributes(user), nil
}

func LoginUser(input AuthRequest) (string, error) {

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

func buildUserAtributes(user models.User) UserResponse {
	return UserResponse{
		Id:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}
