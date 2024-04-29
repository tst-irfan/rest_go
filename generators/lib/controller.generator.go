package lib

import (
	"fmt"
	"path/filepath"
	"strings"
)

const controllerDir = "app/controllers"

func GenerateController(input FileGenerator) error {
	controllerPath := filepath.Join(controllerDir, fmt.Sprintf("%s.controller.go", strings.ToLower(input.Name)))
	controllerName := strings.Title(input.Name)
	content := ControllerTemplate(controllerName, strings.ToLower(input.Name))
	err := createFile(controllerPath, content)
	if err != nil {
		return err
	}
	return nil
}

func ControllerTemplate(controllerName string, controllerNameLower string) string {
	content := `package controllers

	import (
		"net/http"
		"rest_go/app/helpers"
		"rest_go/app/services"
		"rest_go/app/types"
		"rest_go/app/models"
		"strconv"
	
		"github.com/gin-gonic/gin"
	)
	
	func GetAll%Ss(c *gin.Context) {
		var %ss []models.%S
		var err error
		var metadata types.MetaData
	
		%ss, err, metadata = services.ShowAll%S()
	
		if err != nil {
			helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
			return
		}
	
		helpers.ResponseSuccessWithMeta(c, "%S found", %ss, http.StatusOK, metadata)
	}

	func Get%S(c *gin.Context) {
		IDStr := c.Param("id")
		ID, err := strconv.ParseUint(IDStr, 10, 64)
		if err != nil {
			helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
			return
		}
	
		%s, err := services.Get%SByID(uint(ID))
	
		if err != nil {
			helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
			return
		}
	
		helpers.ResponseSuccess(c, "%S found", %s, http.StatusOK)
	}

	func Create%S(c *gin.Context) {
		var input types.%SRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
			return
		}
	
		%s, err := services.Create%S(input)

		if err != nil {
			helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
			return
		}
	
		helpers.ResponseSuccess(c, "%S created", %s, http.StatusCreated)
	}

	func Update%S(c *gin.Context) {
		IDStr := c.Param("id")
		ID, err := strconv.ParseUint(IDStr, 10, 64)
		if err != nil {
			helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
			return
		}
	
		var input types.%SRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
			return
		}
	
		%s, err := services.Update%S(uint(ID), input)
	
		if err != nil {
			helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
			return
		}
	
		helpers.ResponseSuccess(c, "%S updated", %s, http.StatusOK)
	}

	func Delete%S(c *gin.Context) {
		IDStr := c.Param("id")
		ID, err := strconv.ParseUint(IDStr, 10, 64)
		if err != nil {
			helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
			return
		}
	
		err = services.Delete%S(uint(ID))
		if err != nil {
			helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
			return
		}
	
		helpers.ResponseSuccess(c, "%S deleted", "", http.StatusOK)
	}`
	content = strings.Replace(content, "%S", controllerName, -1)
	content = strings.Replace(content, "%s", controllerNameLower, -1)
	return content
}
