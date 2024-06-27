package lib

import (
	"fmt"
	"path/filepath"
	"strings"

	"rest_go/generators/templates"
	"rest_go/generators/types"
)

const (
	controllerDir  = "app/controllers"
	modelDir       = "app/models"
	serviceDir     = "app/services"
	routerDir      = "app/routers"
	autoMigratePath = "initializer/auto.migrate.go"
	requestDir     = "app/types"
	routerPath     = "app/routers/router.go"
)

func Generate(name string, generator string, fieldArgs []types.FieldArgs) error {
	generatorMap := map[string]struct {
		dir     string
		format  string
		content func(string, []types.FieldArgs) string
		after   func(string) error
	}{
		"controller": {controllerDir, "%s.controller.go", func(n string, _ []types.FieldArgs) string {
			return templates.ControllerTemplate(n)
		}, nil},
		"model": {modelDir, "%s.go", templates.ModelTemplate, AppendAutoMigrate},
		"service": {serviceDir, "%s.service.go", func(n string, _ []types.FieldArgs) string {
			return templates.ServiceTemplate(n)
		}, nil},
		"request": {requestDir, "%s.request.go", templates.RequestTypeTemplate, nil},
		"router": {routerDir, "%s.router.go", func(n string, _ []types.FieldArgs) string {
			return templates.RouterTemplate(n)
		}, AppendRouter},
	}

	gen, ok := generatorMap[generator]
	if !ok {
		return fmt.Errorf("invalid generator type")
	}

	name = strings.Title(name)
	path := filepath.Join(gen.dir, fmt.Sprintf(gen.format, strings.ToLower(name)))
	content := gen.content(name, fieldArgs)

	if err := CreateFile(path, content); err != nil {
		return err
	}
	if gen.after != nil {
		if err := gen.after(name); err != nil {
			return err
		}
	}
	return nil
}

func AppendAutoMigrate(modelName string) error {
	autoMigrateContent := fmt.Sprintf(`
		&models.%s{},`, modelName)
	return AppendContent(autoMigratePath, autoMigrateContent, "model")
}

func AppendRouter(routerName string) error {
	routerContent := fmt.Sprintf(`
	Setup%sRoutes(public)`, routerName)
	return AppendContent(routerPath, routerContent, "router")
}
