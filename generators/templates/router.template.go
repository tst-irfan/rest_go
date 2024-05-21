package templates

import (
	"strings"
	"text/template"
)

func RouterTemplate(routerName string) string {
	content := template.Must(template.New("router").Parse(
		`package routers

import (
	"rest_go/app/controllers"
	"github.com/gin-gonic/gin"
)

func Setup{{.RouterName}}Routes(r *gin.RouterGroup) {
	r.GET("/{{.RouterNameLower}}s", controllers.GetAll{{.RouterName}}s)
	r.GET("/{{.RouterNameLower}}s/:id", controllers.Get{{.RouterName}})
	r.POST("/{{.RouterNameLower}}s", controllers.Create{{.RouterName}})
	r.PUT("/{{.RouterNameLower}}s/:id", controllers.Update{{.RouterName}})
	r.DELETE("/{{.RouterNameLower}}s/:id", controllers.Delete{{.RouterName}})
}

`))
	data := struct {
		RouterName      string
		RouterNameLower string
	}{
		RouterName:      routerName,
		RouterNameLower: strings.ToLower(routerName),
	}
	var buf strings.Builder
	content.Execute(&buf, data)
	return buf.String()
}
