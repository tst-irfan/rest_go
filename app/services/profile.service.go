package services

import (
	"rest_go/app/models"
	"rest_go/app/types"
)

func ShowAllProfiles() ([]types.Profile, error) {
	profiles, err := models.ProfileQuery.FindAll()
	if err != nil {
		return []types.Profile{}, err
	}

	return models.BuildProfilesAttributes(profiles), nil
}

func CreateProfile(input types.ProfileRequest, userID uint) (types.Profile, error) {
	profile := models.Profile{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		DoB:       input.DoB,
		UserID:    userID,
	}

	createdProfile, err := models.ProfileQuery.Create(profile)
	if err != nil {
		return types.Profile{}, err
	}

	return models.BuildProfileAttributes(createdProfile), nil
}

func GetProfileByID(ID uint) (types.Profile, error) {
	profile, err := models.ProfileQuery.FindByID(ID)
	if err != nil {
		return types.Profile{}, err
	}

	return models.BuildProfileAttributes(profile), nil
}

func GetProfileByUserID(userID uint) (types.Profile, error) {
	profile, err := models.ProfileQuery.FindOneByColumn("user_id", userID)
	if err != nil {
		return types.Profile{}, err
	}
	return models.BuildProfileAttributes(profile), nil
}

func UpdateProfile(input types.ProfileRequest, ID uint) (types.Profile, error) {
	profile, err := models.ProfileQuery.FindOneByColumn("user_id", ID)
	if err != nil {
		return types.Profile{}, err
	}

	profile.FirstName = input.FirstName
	profile.LastName = input.LastName
	profile.DoB = input.DoB

	updatedProfile, err := models.ProfileQuery.Update(*profile)
	if err != nil {
		return types.Profile{}, err
	}

	return models.BuildProfileAttributes(updatedProfile), nil
}

func DeleteProfile(userID uint) error {
	profile, err :=  models.ProfileQuery.FindOneByColumn("user_id", userID)
	if err != nil {
		return err
	}

	err = models.ProfileQuery.DeleteByID(profile.ID)
	if err != nil {
		return err
	}

	return nil
}
