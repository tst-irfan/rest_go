package controllers

import (
	"auth_go/app/helpers"
	"auth_go/app/services"
	"auth_go/app/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllProfiles(c *gin.Context) {
	profiles, err := services.ShowAllProfiles()
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "Profiles found", profiles, http.StatusOK)
}

func GetProfile(c *gin.Context) {
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	profile, err := services.GetProfileByID(uint(ID))

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "Profile found", profile, http.StatusOK)
}

func ShowMyProfile(c *gin.Context) {
	userID := c.MustGet("UserID").(uint)

	profile, err := services.GetProfileByUserID(userID)

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "Profile found", profile, http.StatusOK)
}

func SaveProfile(c *gin.Context) {
	var input types.ProfileRequest
	userID := c.MustGet("UserID").(uint)

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	profile, err := services.CreateProfile(input, userID)

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "Profile has been created", profile, http.StatusCreated)

}

func UpdateProfile(c *gin.Context) {
	var input types.ProfileRequest

	ID := c.MustGet("UserID").(uint)

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	profile, err := services.UpdateProfile(input, uint(ID))

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	
	helpers.ResponseSuccess(c, "Profile has been updated", profile, http.StatusOK)
}

func DeleteProfile(c *gin.Context) {

	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.DeleteProfile(uint(ID))
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "Profile has been deleted", "", http.StatusOK)
}
