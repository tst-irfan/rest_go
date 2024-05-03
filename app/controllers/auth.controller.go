package controllers

import (
	"net/http"

	"rest_go/app/helpers"
	"rest_go/app/services"
	. "rest_go/app/types"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body AuthRequest true "User input"
// @Success 201 {object} UserResponse
// @Router /register [post]
func Register(c *gin.Context) {
	var input AuthRequest

	if err := c.ShouldBindJSON(&input); err != nil {

		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	println(input.Email)
	user, err := services.RegisterUser(input)

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "User has been created", user, http.StatusCreated)
}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body AuthRequest true "User input"
// @Success 200 {object} LoginResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var input AuthRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := services.LoginUser(input)

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "User has been logged in", LoginResponse{Token: token}, http.StatusOK)
}
