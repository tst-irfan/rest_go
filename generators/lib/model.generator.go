package lib

import (
	"fmt"
	"path/filepath"
	"rest_go/generators/templates"
	"rest_go/generators/types"
	"strings"
)

const modelDir = "app/models"
const autoMigratePath = "initializer/auto.migrate.go"

func GenerateModel(name string, fieldArgs []types.FieldArgs) error {
	modelPath := filepath.Join(modelDir, fmt.Sprintf("%s.go", strings.ToLower(name)))
	modelName := strings.Title(name)
	content := templates.ModelTemplate(modelName, fieldArgs)
	err := CreateFile(modelPath, content)
	AppendAutoMigrate(modelName)
	if err != nil {
		return err
	}
	return nil
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
