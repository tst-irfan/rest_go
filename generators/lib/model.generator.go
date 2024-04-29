package lib

import (
	"fmt"
	"path/filepath"
	"strings"
)

const modelDir = "app/models"
const typeDir = "app/types"

func GenerateModel(input FileGenerator) error {
	modelPath := filepath.Join(modelDir, fmt.Sprintf("%s.go", strings.ToLower(input.Name)))
	modelName := strings.Title(input.Name)
	content := ModelTemplate(modelName)
	err := createFile(modelPath, content)
	GenerateTypeFiles(modelName)
	if err != nil {
		return err
	}
	return nil
}

func ModelTemplate(modelName string) string {
	return fmt.Sprintf(`package models

import (
	"github.com/jinzhu/gorm"
	"rest_go/db"
	"rest_go/app/helpers"
)

type %s struct {
	gorm.Model
}

var %sQuery = db.QueryHelper[%s]{}
var %sValidation = helpers.ValidationHelper{
	RequiredFields:          []string{},
	ShouldGreaterThanFields: []helpers.Field{},
	ShouldLessThanFields:    []helpers.Field{},
}

`, modelName, modelName, modelName, modelName)
}

func GenerateTypeFiles(modelName string) {
	typeRequestFileName := fmt.Sprintf("%s.type.go", strings.ToLower(modelName)+".request")
	typeResponseFileName := fmt.Sprintf("%s.type.go", strings.ToLower(modelName)+".response")
	GenerateTypeFile(modelName+"Request", typeRequestFileName)
	GenerateTypeFile(modelName+"Response", typeResponseFileName)
}

func GenerateTypeFile(modelName string, fileName string) {
	typePath := filepath.Join(typeDir, fileName)
	content := TypeTemplate(modelName)
	createFile(typePath, content)
}

func TypeTemplate(modelName string) string {
	return fmt.Sprintf(`package types

type %s struct {
}

`, modelName)
}
