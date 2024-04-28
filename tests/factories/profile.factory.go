package factories

import (
	"rest_go/app/models"

	"github.com/brianvoe/gofakeit/v7"
)

func GenerateProfile() *models.Profile {
	user := GenerateUser()
	profileParams := models.Profile{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		DoB:       gofakeit.Date(),
		UserID:    user.ID,
	}
	profile, err := models.ProfileQuery.Create(profileParams)
	if err != nil {
		panic(err)
	}
	return profile
}

func GenerateProfiles(quantity int) []*models.Profile {
	profiles := make([]*models.Profile, quantity)
	for i := 0; i < quantity; i++ {
		profiles[i] = GenerateProfile()
	}
	return profiles
}
