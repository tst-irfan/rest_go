package templates

import (
	"rest_go/generators/types"
	"strings"
	"text/template"
)

func RequestTypeTemplate(name string, fields []types.FieldArgs) string {
	content := template.Must(template.New("requestType").Parse(`
package types

type {{.Name}}Request struct {
{{.Fields}}
}
`))
	data := struct {
		Name   string
		Fields string
	}{
		Name:   name,
		Fields: FieldsTemplate(fields),
	}
	var buf strings.Builder
	content.Execute(&buf, data)
	return buf.String()
}