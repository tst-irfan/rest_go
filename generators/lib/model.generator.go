package lib

import (
	"fmt"
	"path/filepath"
	"strings"
)

const modelDir = "app/models"
const autoMigratePath = "initializer/auto.migrate.go"

func GenerateModel(input FileGenerator) error {
	modelPath := filepath.Join(modelDir, fmt.Sprintf("%s.go", strings.ToLower(input.Name)))
	modelName := strings.Title(input.Name)
	content := ModelTemplate(modelName)
	err := createFile(modelPath, content)
	AppendAutoMigrate(modelName)
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

func (m *%s) BeforeSave() error {
	_, err := %sValidation.Validate(m)
	if err != nil {
		return err
	}
	return nil
}

`, modelName, modelName, modelName, modelName, modelName, modelName)
}

func AppendAutoMigrate(modelName string) error {
	autoMigrateContent := fmt.Sprintf(`
		&models.%s{},`, modelName)
	err := AppendContent(autoMigratePath, autoMigrateContent, "model")
	if err != nil {
		return err
	}
	return nil
}
