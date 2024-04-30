package lib

import (
	"fmt"
	"path/filepath"
	"strings"
)

const routerDir = "app/routers"
const routerPath = "app/routers/router.go"

func GenerateRouter(input FileGenerator) error {
	routerPath := filepath.Join(routerDir, fmt.Sprintf("setup_%s.router.go", strings.ToLower(input.Name)))
	routerName := strings.Title(input.Name)
	content := RouterTemplate(routerName, strings.ToLower(input.Name))
	err := createFile(routerPath, content)
	err = AppendRouter(routerName)
	if err != nil {
		return err
	}
	return nil
}

func RouterTemplate(routerName string, routerNameLower string) string {
	content := `package routers

import (
	"rest_go/app/controllers"
	"github.com/gin-gonic/gin"
)

func Setup%SRoutes(r *gin.RouterGroup) {
	r.GET("/%ss", controllers.GetAll%Ss)
	r.GET("/%ss/:id", controllers.Get%S)
	r.POST("/%ss", controllers.Create%S)
	r.PUT("/%ss/:id", controllers.Update%S)
	r.DELETE("/%ss/:id", controllers.Delete%S)
}

`
	content = strings.Replace(content, "%S", routerName, -1)
	content = strings.Replace(content, "%s", routerNameLower, -1)
	return content
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
