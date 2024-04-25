package models

import (
	"errors"
	"time"

	"auth_go/app/types"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	FirstName string    `gorm:"size:255;not null;" json:"first_name"`
	LastName  string    `gorm:"size:255;not null;" json:"last_name"`
	DoB       time.Time `json:"dob"`
}

func GetAllProfiles() ([]Profile, error) {
	var profiles []Profile
	err := DB.Find(&profiles).Error
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func GetProfileByID(ID uint) (*Profile, error) {
	var profile Profile
	err := DB.Where("id = ?", ID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (profile *Profile) SaveProfile() (*Profile, error) {
	err := DB.Create(&profile).Error
	if err != nil {
		return &Profile{}, err
	}
	return profile, nil
}

func (profile *Profile) UpdateProfile() (*Profile, error) {
	err := DB.Save(&profile).Error
	if err != nil {
		return &Profile{}, err
	}
	return profile, nil
}

func (profile *Profile) DeleteProfile() error {
	err := DB.Delete(&profile).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProfileByUserID(userID uint) (*Profile, error) {
	var profile Profile
	err := DB.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func GetUser(userID uint) User {
	var user User
	DB.Where("id = ?", userID).First(&user)
	return user
}

func (profile *Profile) GetUser() User {
	var user User
	DB.Where("id = ?", profile.UserID).First(&user)
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
	var user User
	DB.Where("id = ?", userId).First(&user)
	return user.ID != 0
}

func checkProfileExists(userId uint) bool {
	var profile Profile
	DB.Where("user_id = ?", userId).First(&profile)
	return profile.UserID != 0
}

func (p *Profile) BeforeUpdate() error {
	if !checkProfileExists(p.UserID) {
		return errors.New("Profile does not exist")
	}
	return nil
}

func BuildProfileAttributes(profile *Profile) types.Profile {
	user := GetUser(profile.UserID)
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
