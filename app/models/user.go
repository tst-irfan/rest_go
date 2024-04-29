package models

import (
	"html"
	"strings"

	"rest_go/app/helpers"
	"rest_go/app/types"
	"rest_go/db"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

var UserQuery = db.QueryHelper[User]{}
var userValidation = helpers.ValidationHelper{
	RequiredFields:          []string{"Email", "Password"},
	ShouldGreaterThanFields: []helpers.Field{},
	ShouldLessThanFields:    []helpers.Field{},
}

func VerifyPassword(providedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *User) BeforeSave() error {
	_, err := UserQuery.FindOneByColumn("email", u.Email)
	if err == nil {
		return errors.New("Email already exists")
	}

	_, err = userValidation.Validate(u)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}

func BuildUserAtributes(user *User) types.UserResponse {
	return types.UserResponse{
		Id:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}
