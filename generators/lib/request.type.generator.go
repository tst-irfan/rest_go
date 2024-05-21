package lib

import (
	"fmt"
	"path/filepath"
	"rest_go/generators/templates"
	"strings"
	"rest_go/generators/types"
)

const requestDir = "app/types"

func GenerateRequest(name string, fieldArgs []types.FieldArgs) error {
	requestPath := filepath.Join(requestDir, fmt.Sprintf("%s.request.go", strings.ToLower(name)))
	requestName := strings.Title(name)
	content := templates.RequestTypeTemplate(requestName, fieldArgs)
	err := createFile(requestPath, content)
	if err != nil {
		return err
	}
	return nil
}