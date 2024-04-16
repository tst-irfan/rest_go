package models

import (
	"auth_go/app/utils/token"
	"html"
	"strings"

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
	existingUser, _ := GetUserByEmail(user.Email)

	if existingUser.Email != "" {
		return &User{}, errors.New("Email already exists")
	}

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

func LoginCheck(email string, password string) (string, error) {
	var err error
	var isPasswordMatch bool

	user, err := GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	isPasswordMatch, err = VerifyPassword(password, user.Password)

	if !isPasswordMatch {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *User) BeforeSave() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}
