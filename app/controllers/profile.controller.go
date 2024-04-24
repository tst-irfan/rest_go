package controllers

import (
	"auth_go/app/services"
	"auth_go/app/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllProfiles(c *gin.Context) {
	profiles, err := services.ShowAllProfiles()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, profiles)
}

func GetProfile(c *gin.Context) {
	var response types.ProfileResponse

	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		response = types.ProfileResponse{
			Error:   err.Error(),
			Profile: types.Profile{},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	profile, err := services.GetProfileByID(uint(ID))

	if err != nil {
		response = types.ProfileResponse{
			Error:   err.Error(),
			Profile: types.Profile{},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = types.ProfileResponse{
		Error:   "",
		Profile: profile,
	}
	c.JSON(http.StatusOK, response)
}

func ShowMyProfile(c *gin.Context) {
	var response types.ProfileResponse
	userID := c.MustGet("UserID").(uint)

	profile, err := services.GetProfileByUserID(userID)

	if err != nil {
		response = types.ProfileResponse{
			Error:   err.Error(),
			Profile: types.Profile{},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = types.ProfileResponse{
		Error:   "",
		Profile: profile,
	}
	c.JSON(http.StatusOK, response)
}

func SaveProfile(c *gin.Context) {
	var input types.ProfileRequest
	var response types.ProfileResponse
	userID := c.MustGet("UserID").(uint)

	if err := c.ShouldBindJSON(&input); err != nil {
		response = types.ProfileResponse{
			Error:   err.Error(),
			Profile: types.Profile{},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	profile, err := services.CreateProfile(input, userID)

	if err != nil {
		response = types.ProfileResponse{
			Error:   err.Error(),
			Profile: types.Profile{},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = types.ProfileResponse{
		Error:   "",
		Profile: profile,
	}
	c.JSON(http.StatusOK, response)

}



func UpdateProfile(c *gin.Context) {
	var input types.ProfileRequest
	var response types.ProfileResponse

	ID := c.MustGet("UserID").(uint)

	if err := c.ShouldBindJSON(&input); err != nil {
		response = types.ProfileResponse{
			Error:   err.Error(),
			Profile: types.Profile{},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	profile, err := services.UpdateProfile(input, uint(ID))

	if err != nil {
		response = types.ProfileResponse{
			Error:   err.Error(),
			Profile: types.Profile{},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = types.ProfileResponse{
		Error:   "",
		Profile: profile,
	}
	c.JSON(http.StatusOK, response)

}

func DeleteProfile(c *gin.Context) {

	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = services.DeleteProfile(uint(ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Profile deleted successfully")
}
