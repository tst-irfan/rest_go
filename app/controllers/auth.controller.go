package controllers

import (
	"net/http"

	"auth_go/app/helpers"
	"auth_go/app/services"
	. "auth_go/app/types"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input AuthRequest

	if err := c.ShouldBindJSON(&input); err != nil {

		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := services.RegisterUser(input)

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "User has been created", user, http.StatusCreated)
}

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
