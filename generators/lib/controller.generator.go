package lib

import (
	"fmt"
	"path/filepath"
	"rest_go/generators/templates"
	"strings"
)

const controllerDir = "app/controllers"

func GenerateController(name string) error {
	controllerPath := filepath.Join(controllerDir, fmt.Sprintf("%s.controller.go", strings.ToLower(name)))
	controllerName := strings.Title(name)
	content := templates.ControllerTemplate(controllerName)
	err := createFile(controllerPath, content)
	if err != nil {
		return err
	}
	return nil
}
