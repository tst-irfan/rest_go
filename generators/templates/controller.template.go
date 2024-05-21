package templates

import (
	"strings"
	"text/template"
)

func ControllerTemplate(controllerName string) string {
	content := template.Must(template.New("controller").Parse(`
package controllers

import (
	"net/http"
	"rest_go/app/helpers"
	"rest_go/app/services"
	"rest_go/app/types"
	"rest_go/app/models"
	"rest_go/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Show All {{.ControllerName}}s godoc
// @Summary Show all {{.ControllerNameLower}}s
// @Description get all {{.ControllerNameLower}}s
// @Tags {{.ControllerName}}
// @Accept  json
// @Produce  json
// @Success 200 {object} types.SuccessWithMeta[[]models.{{.ControllerName}}]
// @Failure 400 {object} types.Error
// @Router /{{.ControllerNameLower}}s [get]
// @Security Bearer
func GetAll{{.ControllerName}}s(c *gin.Context) {
	var {{.ControllerNameLower}}s []models.{{.ControllerName}}
	var err error
	var metadata types.MetaData

	{{.ControllerNameLower}}s, err, metadata = services.ShowAll{{.ControllerName}}()

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccessWithMeta(c, "{{.ControllerName}} found", {{.ControllerNameLower}}s, http.StatusOK, metadata)
}

// Get {{.ControllerName}} godoc
// @Summary Get a {{.ControllerNameLower}}
// @Description get a {{.ControllerNameLower}} by id
// @Tags {{.ControllerName}}
// @Accept  json
// @Produce  json
// @Param id path int true "{{.ControllerName}} ID"
// @Success 200 {object} types.Success[models.{{.ControllerName}}]
// @Failure 400 {object} types.Error
// @Router /{{.ControllerNameLower}}s/{id} [get]
// @Security Bearer
func Get{{.ControllerName}}(c *gin.Context) {
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	{{.ControllerNameLower}}, err := services.Get{{.ControllerName}}ByID(uint(ID))

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "{{.ControllerName}} found", {{.ControllerNameLower}}, http.StatusOK)
}

// Create {{.ControllerName}} godoc
// @Summary Create a {{.ControllerNameLower}}
// @Description create a {{.ControllerNameLower}}
// @Tags {{.ControllerName}}
// @Accept  json
// @Produce  json
// @Param input body types.{{.ControllerName}}Request true "User input"
// @Success 201 {object} types.Success[models.{{.ControllerName}}]
// @Failure 400 {object} types.Error
// @Router /{{.ControllerNameLower}}s [post]
// @Security Bearer
func Create{{.ControllerName}}(c *gin.Context) {
	var input types.{{.ControllerName}}Request
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	{{.ControllerNameLower}}Params, err := utils.TypeConverter[models.{{.ControllerName}}](&input)

	{{.ControllerNameLower}}, err := services.Create{{.ControllerName}}(*{{.ControllerNameLower}}Params)

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "{{.ControllerName}} created", {{.ControllerNameLower}}, http.StatusCreated)
}

// Update {{.ControllerName}} godoc
// @Summary Update a {{.ControllerNameLower}}
// @Description update a {{.ControllerNameLower}} by id
// @Tags {{.ControllerName}}
// @Accept  json
// @Produce  json
// @Param id path int true "{{.ControllerName}} ID"
// @Param input body types.{{.ControllerName}}Request true "User input"
// @Success 200 {object} types.Success[models.{{.ControllerName}}]
// @Failure 400 {object} types.Error
// @Router /{{.ControllerNameLower}}s/{id} [put]
// @Security Bearer
func Update{{.ControllerName}}(c *gin.Context) {
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var input types.{{.ControllerName}}Request
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	{{.ControllerNameLower}}Params, err := utils.TypeConverter[models.{{.ControllerName}}](&input)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	{{.ControllerNameLower}}, err := services.Update{{.ControllerName}}(uint(ID), *{{.ControllerNameLower}}Params)

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "{{.ControllerName}} updated", {{.ControllerNameLower}}, http.StatusOK)
}

// Delete {{.ControllerName}} godoc
// @Summary Delete a {{.ControllerNameLower}}
// @Description delete a {{.ControllerNameLower}} by id
// @Tags {{.ControllerName}}
// @Accept  json
// @Produce  json
// @Param id path int true "{{.ControllerName}} ID"
// @Success 200 {string} types.Success
// @Failure 400 {object} types.Error
// @Router /{{.ControllerNameLower}}s/{id} [delete]
// @Security Bearer
func Delete{{.ControllerName}}(c *gin.Context) {
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.Delete{{.ControllerName}}(uint(ID))
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "{{.ControllerName}} deleted", "", http.StatusOK)
}
`))
	data := struct {
		ControllerName      string
		ControllerNameLower string
	}{
		ControllerName:      controllerName,
		ControllerNameLower: strings.ToLower(controllerName),
	}
	var buf strings.Builder
	content.Execute(&buf, data)
	return buf.String()
}
