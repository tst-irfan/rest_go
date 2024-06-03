package templates

import (
	"strings"
	"text/template"
)

func ServiceTemplate(serviceName string) string {
	content := template.Must(template.New("service").Parse(`
package services

import (
	"rest_go/app/models"
	"rest_go/app/types"
	"sync"
)

func ShowAll{{.ServiceName}}() ([]models.{{.ServiceName}}, error, types.MetaData) {
	var wg sync.WaitGroup

	var {{.ServiceNameLower}}s []models.{{.ServiceName}}
	var totalItems int
	var err error

	wg.Add(2)
	go func() {
		defer wg.Done()
		{{.ServiceNameLower}}s, err = models.{{.ServiceName}}Query.FindAll()
	}()
	go func() {
		defer wg.Done()
		totalItems, err = models.{{.ServiceName}}Query.Count()
	}()
	
	wg.Wait()

	if err != nil {
		return []models.{{.ServiceName}}{}, err, types.MetaData{}
	}

	metaData := types.MetaData{
		TotalItems: totalItems,
	}

	return {{.ServiceNameLower}}s, nil, metaData
}

func Create{{.ServiceName}}({{.ServiceNameLower}} models.{{.ServiceName}}) (models.{{.ServiceName}}, error) {
	created{{.ServiceName}}, err := models.{{.ServiceName}}Query.Create({{.ServiceNameLower}})
	if err != nil {
		return models.{{.ServiceName}}{}, err
	}

	return *created{{.ServiceName}}, nil
}

func Get{{.ServiceName}}ByID(ID uint) (models.{{.ServiceName}}, error) {
	{{.ServiceNameLower}}, err := models.{{.ServiceName}}Query.FindByID(ID)
	if err != nil {
		return models.{{.ServiceName}}{}, err
	}

	return *{{.ServiceNameLower}}, nil
}

func Update{{.ServiceName}}(ID uint, {{.ServiceNameLower}}Params models.{{.ServiceName}}) (models.{{.ServiceName}}, error) {
	{{.ServiceNameLower}}, err := models.{{.ServiceName}}Query.FindByID(ID)
	if err != nil {
		return models.{{.ServiceName}}{}, err
	}

	{{.ServiceNameLower}}Params.ID = {{.ServiceNameLower}}.ID

	updated{{.ServiceName}}, err := models.{{.ServiceName}}Query.Update({{.ServiceNameLower}}Params)
	if err != nil {
		return models.{{.ServiceName}}{}, err
	}

	return *updated{{.ServiceName}}, nil
}

func Delete{{.ServiceName}}(ID uint) error {
	err := models.{{.ServiceName}}Query.DeleteByID(ID)
	if err != nil {
		return err
	}

	return nil
}
`))
	data := struct {
		ServiceName      string
		ServiceNameLower string
	}{
		ServiceName:      serviceName,
		ServiceNameLower: strings.ToLower(serviceName),
	}
	var buf strings.Builder
	content.Execute(&buf, data)
	return buf.String()
}
