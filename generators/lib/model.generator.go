package lib

import (
	"fmt"
	"path/filepath"
	"strings"
)

const modelDir = "app/models"
const autoMigratePath = "initializer/auto.migrate.go"

func GenerateModel(input FileGenerator, fieldArgs []FieldArgs) error {
	modelPath := filepath.Join(modelDir, fmt.Sprintf("%s.go", strings.ToLower(input.Name)))
	modelName := strings.Title(input.Name)
	content := ModelTemplate(modelName, fieldArgs)
	err := createFile(modelPath, content)
	AppendAutoMigrate(modelName)
	if err != nil {
		return err
	}
	return nil
}

func ModelTemplate(modelName string, fieldArgs []FieldArgs) string {
	fields := FieldsTemplate(fieldArgs)
	return fmt.Sprintf(`package models

import (
	"github.com/jinzhu/gorm"
	"rest_go/db"
	"rest_go/app/helpers"
)

type %s struct {
	gorm.Model
%s
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

`, modelName, fields, modelName, modelName, modelName, modelName, modelName)
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

func FieldsTemplate(field []FieldArgs) string {
	var fields string
	for _, f := range field {
		fields += fmt.Sprintf("\t%s %s `json:\"%s`\n", strings.Title(f.Name), f.Type, toSnakeCase(f.Name))
	}
	return fields
}

func toSnakeCase(s string) string {
	var snake string
	for i, c := range s {
		if 'A' <= c && c <= 'Z' {
			if i > 0 {
				snake += "_"
			}
			snake += string(c - 'A' + 'a')
		} else {
			snake += string(c)
		}
	}
	return snake
}
