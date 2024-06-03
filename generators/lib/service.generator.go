package lib

import (
	"fmt"
	"path/filepath"
	"rest_go/generators/templates"
	"strings"
)

const serviceDir = "app/services"

func GenerateService(name string) error {
	servicePath := filepath.Join(serviceDir, fmt.Sprintf("%s.service.go", strings.ToLower(name)))
	serviceName := strings.Title(name)
	content := templates.ServiceTemplate(serviceName)
	err := CreateFile(servicePath, content)
	if err != nil {
		return err
	}
	return nil
}
