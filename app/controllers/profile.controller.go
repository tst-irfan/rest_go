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

// Show All Profiles godoc
// @Summary Show all profiles
// @Description get all profiles
// @Tags profile
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Success 200 {object} []types.Profile
// @Router /profiles [get]
// @Security Bearer
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

// Get Profile godoc
// @Summary Get a profile
// @Description get a profile by id
// @Tags profile
// @Accept  json
// @Produce  json
// @Param id path int true "Profile ID"
// @Success 200 {object} types.Profile
// @Router /profiles/{id} [get]
// @Security Bearer
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

// Show My Profile godoc
// @Summary Show my profile
// @Description get my profile
// @Tags profile
// @Accept  json
// @Produce  json
// @Success 200 {object} types.Profile
// @Router /my-profile [get]
// @Security Bearer
func ShowMyProfile(c *gin.Context) {
	userID := c.MustGet("UserID").(uint)

	profile, err := services.GetProfileByUserID(userID)

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "Profile found", profile, http.StatusOK)
}

// Save Profile godoc
// @Summary Save a profile
// @Description Save a profile
// @Tags profile
// @Accept  json
// @Produce  json
// @Param input body types.ProfileRequest true "Profile input"
// @Success 201 {object} types.Profile
// @Router /profiles [post]
// @Security Bearer
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

// Update Profile godoc
// @Summary Update a profile
// @Description Update a profile
// @Tags profile
// @Accept  json
// @Produce  json
// @Param input body types.ProfileRequest true "Profile input"
// @Success 200 {object} types.Profile
// @Router /profiles [put]
// @Security Bearer
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

// Delete Profile godoc
// @Summary Delete a profile
// @Description Delete a profile
// @Tags profile
// @Accept  json
// @Produce  json
// @Param id path int true "Profile ID"
// @Success 200 {string} string
// @Router /profiles/{id} [delete]
// @Security Bearer
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
