package templates

import (
	"rest_go/generators/types"
	"strings"
	"text/template"
)


func ModelTemplate(modelName string, fieldArgs []types.FieldArgs) string {
	content := template.Must(template.New("model").Parse(`
package models
import (
	"github.com/jinzhu/gorm"
	"rest_go/db"
	"rest_go/app/helpers"
)

type {{.ModelName}} struct {
	gorm.Model
{{.Fields}}
}

var {{.ModelName}}Query = db.QueryHelper[{{.ModelName}}]{}
var {{.ModelName}}Validation = helpers.ValidationHelper{
	RequiredFields:          []string{},
	ShouldGreaterThanFields: []helpers.Field{},
	ShouldLessThanFields:    []helpers.Field{},
}

func (m *{{.ModelName}}) BeforeSave() error {
	_, err := {{.ModelName}}Validation.Validate(m)
	if err != nil {
		return err
	}
	return nil
}
`))
	data := struct {
		ModelName string
		Fields    string
	}{
		ModelName: modelName,
		Fields:    FieldsTemplate(fieldArgs),
	}
	var buf strings.Builder
	content.Execute(&buf, data)
	return buf.String()
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
