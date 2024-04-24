package services

import (
	"auth_go/app/models"
	"auth_go/app/types"
)

func ShowAllProfiles() ([]types.Profile, error) {
	profiles, err := models.GetAllProfiles()
	if err != nil {
		return []types.Profile{}, err
	}

	return models.BuildProfilesAttributes(profiles), nil
}

func CreateProfile(input types.ProfileRequest) (types.Profile, error) {
	profile := models.Profile{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		DoB:       input.DoB,
		UserID:    input.UserId,
	}

	createdProfile, err := profile.SaveProfile()
	if err != nil {
		return types.Profile{}, err
	}

	return models.BuildProfileAttributes(createdProfile), nil
}

func GetProfileByID(userID uint) (types.Profile, error) {
	profile, err := models.GetProfileByID(userID)
	if err != nil {
		return types.Profile{}, err
	}

	return models.BuildProfileAttributes(profile), nil
}

func UpdateProfile(input types.ProfileRequest, ID uint) (types.Profile, error) {
	profile, err := models.GetProfileByID(ID)
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
