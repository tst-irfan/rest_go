package controllers

import (
	"net/http"

	"auth_go/app/services"
	. "auth_go/app/types"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input AuthRequest
	var response RegisterResponse

	if err := c.ShouldBindJSON(&input); err != nil {
		response = RegisterResponse{
			Error: err.Error(),
			User:  UserResponse{},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := services.RegisterUser(input)

	if err != nil {
		response = RegisterResponse{
			Error: err.Error(),
			User:  UserResponse{},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = RegisterResponse{
		Error: "",
		User:  user,
	}
	c.JSON(http.StatusOK, response)

}

func Login(c *gin.Context) {
	var input AuthRequest
	var response LoginResponse
	var status int

	if err := c.ShouldBindJSON(&input); err != nil {
		response = LoginResponse{
			Error: err.Error(),
			Token: "",
		}
		status = http.StatusBadRequest
		c.JSON(status, response)
		return
	}

	token, err := services.LoginUser(input)

	if err != nil {
		response = LoginResponse{
			Error: err.Error(),
			Token: "",
		}
	} else {
		response = LoginResponse{
			Error: "",
			Token: token,
		}
	}

	c.JSON(status, response)
}
