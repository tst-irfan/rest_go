package lib

import (
	"fmt"
	"path/filepath"
	"rest_go/generators/templates"
	"strings"
)

const routerDir = "app/routers"
const routerPath = "app/routers/router.go"

func GenerateRouter(name string) error {
	routerPath := filepath.Join(routerDir, fmt.Sprintf("setup_%s.router.go", strings.ToLower(name)))
	routerName := strings.Title(name)
	content := templates.RouterTemplate(routerName)
	err := CreateFile(routerPath, content)
	if err != nil {
		return err
	}
	err = AppendRouter(routerName)
	if err != nil {
		return err
	}
	return nil
}

func AppendRouter(routerName string) error {
	routerContent := fmt.Sprintf(`
	Setup%sRoutes(public)
`, routerName)
	err := AppendContent(routerPath, routerContent, "router")
	if err != nil {
		return err
	}
	return nil
}
