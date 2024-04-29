package controllers

import (
	"fmt"
	"net/http"
	"rest_go/app/helpers"
	"rest_go/app/services"
	"rest_go/app/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllProfiles(c *gin.Context) {
	var input types.GetProfilesRequest
	var profiles []types.Profile
	var err error
	var metadata types.MetaData

	if err := c.ShouldBindQuery(&input); err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("input: %+v", input)
	if input.Size == 0 {
		profiles, err, metadata = services.ShowAllProfiles()
	} else {
		profiles, err, metadata = services.ShowAllProfilesWithPagination(input.Page, input.Size)
	}

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccessWithMeta(c, "Profiles found", profiles, http.StatusOK, metadata)
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
