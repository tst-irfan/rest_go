package models

import (
	"errors"
	"time"

	"rest_go/app/types"
	"rest_go/db"

	"rest_go/app/helpers"
	"strconv"

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
var profileValidation = helpers.ValidationHelper{
	RequiredFields:          []string{"FirstName", "LastName", "DoB", "UserID"},
	ShouldGreaterThanFields: []helpers.Field{},
	ShouldLessThanFields: []helpers.Field{
		{
			Name:  "DoB",
			Type:  "time",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		},
	},
}

func (p *Profile) AfterFind() {
	user, err := UserQuery.FindByID(p.UserID)
	if err == nil {
		p.User = *user
	}
}

func (p *Profile) BeforeCreate() error {
	_, err := profileValidation.Validate(p)
	if err != nil {
		return err
	}
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
	_, err := profileValidation.Validate(p)
	if err != nil {
		return err
	}
	if !checkProfileExists(p.UserID) {
		return errors.New("Profile does not exist")
	}
	return nil
}

func BuildProfileAttributes(profile *Profile) types.Profile {
	userAttributes := BuildUserAtributes(&profile.User)
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
