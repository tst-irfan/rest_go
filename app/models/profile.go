package models

import (
	"errors"
	"time"

	"rest_go/app/types"
	"rest_go/db"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	User      User      `json:"user"`
	UserID    uint      `json:"user_id"`
	FirstName string    `gorm:"size:255;not null;" json:"first_name"`
	LastName  string    `gorm:"size:255;not null;" json:"last_name"`
	DoB       time.Time `json:"dob"`
}

var ProfileQuery = db.QueryHelper[Profile]{}

func GetUser(userID uint) *User {
	user, err := UserQuery.FindByID(userID)
	if err != nil {
		return &User{}
	}
	return user
}

func (profile *Profile) GetUser() *User {
	user, err := UserQuery.FindByID(profile.UserID)
	if err != nil {
		return &User{}
	}
	return user
}

func (p *Profile) BeforeCreate() error {
	if checkProfileExists(p.UserID) {
		return errors.New("Profile already exists")
	}
	if !checkUserExists(p.UserID) {
		return errors.New("User does not exist")
	}

	return nil
}

func (p *Profile) BeforeDelete() error {
	if !checkProfileExists(p.UserID) {
		return errors.New("Profile does not exist")
	}
	return nil
}

func checkUserExists(userId uint) bool {
	user, err := UserQuery.FindByID(userId)
	if err != nil {
		return false
	}
	return user.ID != 0
}

func checkProfileExists(userId uint) bool {
	_, err := ProfileQuery.FindOneByColumn("user_id", userId)
	return err == nil
}

func (p *Profile) BeforeUpdate() error {
	if !checkProfileExists(p.UserID) {
		return errors.New("Profile does not exist")
	}
	return nil
}

func BuildProfileAttributes(profile *Profile) types.Profile {
	user := profile.GetUser()
	userAttributes := BuildUserAtributes(user)
	return types.Profile{
		Id:        profile.ID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		DoB:       profile.DoB.String(),
		User:      userAttributes,
	}
}

func BuildProfilesAttributes(profiles []Profile) []types.Profile {
	var profilesAttributes []types.Profile
	for _, profile := range profiles {
		profileAttributes := BuildProfileAttributes(&profile)
		profilesAttributes = append(profilesAttributes, profileAttributes)
	}
	return profilesAttributes
}
