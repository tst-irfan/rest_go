package models

import (
	"html"
	"strings"

	"rest_go/app/types"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (user *User) SaveUser() (*User, error) {

	err := DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func VerifyPassword(providedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *User) BeforeSave() error {
	existingUser, _ := GetUserByEmail(u.Email)

	if existingUser.Email != "" {
		return errors.New("Email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}

func BuildUserAtributes(user User) types.UserResponse {
	return types.UserResponse{
		Id:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}
