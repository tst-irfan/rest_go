package factories

import (
	"rest_go/app/models"

	"github.com/brianvoe/gofakeit/v7"
)

func GenerateProfile() models.Profile {
	var profile models.Profile
	user := GenerateUser()
	profile.UserID = user.ID
	profile.FirstName = gofakeit.FirstName()
	profile.LastName = gofakeit.LastName()
	profile.DoB = gofakeit.Date()

	profile.SaveProfile()
	return profile
}

func GenerateProfiles(n int) []models.Profile {
	profiles := []models.Profile{}
	for i := 0; i < n; i++ {
		profiles = append(profiles, GenerateProfile())
	}
	return profiles
}
