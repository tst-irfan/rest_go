package models_test

import (
	"auth_go/app/models"
	"auth_go/tests"
	"auth_go/tests/factories"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func TestGetAllProfiles(t *testing.T) {
	tests.Setup()
	defer tests.Teardown()

	profiles := factories.GenerateProfiles(10)

	data, err := models.GetAllProfiles()
	if err != nil {
		t.Errorf("Error occured while fetching profiles: %v", err)
	}

	if len(data) != len(profiles) {
		t.Errorf("Expected profiles to be %v, but got %v", len(profiles), len(data))
	}
}

func TestGetProfileByID(t *testing.T) {
	tests.Setup()
	defer tests.Teardown()

	profile := factories.GenerateProfile()

	data, err := models.GetProfileByID(profile.ID)
	if err != nil {
		t.Errorf("Error occured while fetching profile: %v", err)
	}

	if data.ID != profile.ID {
		t.Errorf("Expected profile ID to be %v, but got %v", profile.ID, data.ID)
	}
	if data.FirstName != profile.FirstName {
		t.Errorf("Expected first name to be %v, but got %v", profile.FirstName, data.FirstName)
	}
	if data.LastName != profile.LastName {
		t.Errorf("Expected last name to be %v, but got %v", profile.LastName, data.LastName)
	}
}

func TestSaveProfile(t *testing.T) {
	tests.Setup()
	defer tests.Teardown()

	user := factories.GenerateUser()

	profile := models.Profile{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		DoB:       gofakeit.Date(),
		UserID:    user.ID,
	}

	data, err := profile.SaveProfile()
	if err != nil {
		t.Errorf("Error occured while saving profile: %v", err)
	}

	if data.ID == 0 {
		t.Errorf("Expected profile ID to be greater than 0, but got %v", data.ID)
	}
	if data.FirstName != profile.FirstName {
		t.Errorf("Expected first name to be %v, but got %v", profile.FirstName, data.FirstName)
	}
	if data.LastName != profile.LastName {
		t.Errorf("Expected last name to be %v, but got %v", profile.LastName, data.LastName)
	}

	data, err = profile.SaveProfile()
	if err == nil {
		t.Errorf("Expected error to be thrown")
	}
}

func TestUpdateProfile(t *testing.T) {
	tests.Setup()
	defer tests.Teardown()

	profile := factories.GenerateProfile()

	profile.FirstName = gofakeit.FirstName()
	profile.LastName = gofakeit.LastName()
	profile.DoB = gofakeit.Date()

	data, err := profile.UpdateProfile()
	if err != nil {
		t.Errorf("Error occured while updating profile: %v", err)
	}

	if data.ID != profile.ID {
		t.Errorf("Expected profile ID to be %v, but got %v", profile.ID, data.ID)
	}
	if data.FirstName != profile.FirstName {
		t.Errorf("Expected first name to be %v, but got %v", profile.FirstName, data.FirstName)
	}
	if data.LastName != profile.LastName {
		t.Errorf("Expected last name to be %v, but got %v", profile.LastName, data.LastName)
	}
}

func TestDeleteProfile(t *testing.T) {
	tests.Setup()
	defer tests.Teardown()

	profile := factories.GenerateProfile()

	err := profile.DeleteProfile()
	if err != nil {
		t.Errorf("Error occured while deleting profile: %v", err)
	}

	err = profile.DeleteProfile()
	if err == nil {
		t.Errorf("Expected error to be thrown")
	}

	_, err = models.GetProfileByID(profile.ID)
	if err == nil {
		t.Errorf("Error occured while fetching profile: %v", err)
	}
	err = profile.DeleteProfile()
	if err == nil {
		t.Errorf("Expected error to be thrown")
	}
}

func TestGetProfileByUserID(t *testing.T) {
	tests.Setup()
	defer tests.Teardown()

	profile := factories.GenerateProfile()

	data, err := models.GetProfileByUserID(profile.UserID)
	if err != nil {
		t.Errorf("Error occured while fetching profile: %v", err)
	}

	if data.ID != profile.ID {
		t.Errorf("Expected profile ID to be %v, but got %v", profile.ID, data.ID)
	}
	if data.FirstName != profile.FirstName {
		t.Errorf("Expected first name to be %v, but got %v", profile.FirstName, data.FirstName)
	}
	if data.LastName != profile.LastName {
		t.Errorf("Expected last name to be %v, but got %v", profile.LastName, data.LastName)
	}
}

func TestGetUser(t *testing.T) {
	tests.Setup()
	defer tests.Teardown()

	profile := factories.GenerateProfile()

	data := profile.GetUser()

	if data.ID != profile.UserID {
		t.Errorf("Expected user ID to be %v, but got %v", profile.UserID, data.ID)
	}
}

