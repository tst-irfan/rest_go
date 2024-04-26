package services

import (
	"rest_go/app/models"
	"rest_go/app/types"
)

func ShowAllProfiles() ([]types.Profile, error) {
	profiles, err := models.GetAllProfiles()
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

	createdProfile, err := profile.SaveProfile()
	if err != nil {
		return types.Profile{}, err
	}

	return models.BuildProfileAttributes(createdProfile), nil
}

func GetProfileByID(ID uint) (types.Profile, error) {
	profile, err := models.GetProfileByID(ID)
	if err != nil {
		return types.Profile{}, err
	}

	return models.BuildProfileAttributes(profile), nil
}

func GetProfileByUserID(userID uint) (types.Profile, error) {
	profile, err := models.GetProfileByUserID(userID)
	if err != nil {
		return types.Profile{}, err
	}
	return models.BuildProfileAttributes(profile), nil
}

func UpdateProfile(input types.ProfileRequest, ID uint) (types.Profile, error) {
	profile, err := models.GetProfileByUserID(ID)
	if err != nil {
		return types.Profile{}, err
	}

	profile.FirstName = input.FirstName
	profile.LastName = input.LastName
	profile.DoB = input.DoB

	updatedProfile, err := profile.UpdateProfile()
	if err != nil {
		return types.Profile{}, err
	}

	return models.BuildProfileAttributes(updatedProfile), nil
}

func DeleteProfile(userID uint) error {
	profile, err := models.GetProfileByID(userID)
	if err != nil {
		return err
	}

	err = profile.DeleteProfile()
	if err != nil {
		return err
	}

	return nil
}
