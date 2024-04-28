package models_test

import (
	"rest_go/app/models"
	"rest_go/tests"
	"rest_go/tests/factories"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func TestSaveProfile(t *testing.T) {
	tests.SetupTest()
	user := factories.GenerateUser()
	profileParams := models.Profile{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		DoB:       gofakeit.Date(),
		UserID:    user.ID,
	}
	profile, err := models.ProfileQuery.Create(profileParams)
	if err != nil {
		t.Errorf("Error while saving profile: %v", err)
	}
	if profile.ID == 0 {
		t.Errorf("Error while saving profile: ID is zero")
	}
	if profile.FirstName != profileParams.FirstName {
		t.Errorf("Error while saving profile: FirstName is different")
	}
	if profile.LastName != profileParams.LastName {
		t.Errorf("Error while saving profile: LastName is different")
	}
	if profile.DoB != profileParams.DoB {
		t.Errorf("Error while saving profile: DoB is different")
	}
	if profile.UserID != profileParams.UserID {
		t.Errorf("Error while saving profile: UserID is different")
	}
	tests.TeardownTest()
}
