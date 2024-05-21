package templates

import (
	"fmt"
	"rest_go/generators/types"
	"strings"
)

func FieldsTemplate(field []types.FieldArgs) string {
	var fields string
	for _, f := range field {
		fields += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", strings.Title(f.Name), f.Type, toSnakeCase(f.Name))
	}
	return fields
}