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
	"rest_go/db"
)

type {{.ServiceName}}Service struct {
	Query db.QueryHelperInterface[models.{{.ServiceName}}]
}

type {{.ServiceName}}ServiceInterface interface {
	ShowAll{{.ServiceName}}() ([]models.{{.ServiceName}}, error, types.MetaData)
	Create{{.ServiceName}}(models.{{.ServiceName}}) (models.{{.ServiceName}}, error)
	Get{{.ServiceName}}ByID(uint) (models.{{.ServiceName}}, error)
	Update{{.ServiceName}}(uint, models.{{.ServiceName}}) (models.{{.ServiceName}}, error)
	Delete{{.ServiceName}}(uint) error
}

func New{{.ServiceName}}Service() *{{.ServiceName}}Service {
	return &{{.ServiceName}}Service{Query: &models.{{.ServiceName}}Query}
}

func (s *{{.ServiceName}}Service) ShowAll{{.ServiceName}}() ([]models.{{.ServiceName}}, error, types.MetaData) {
	var wg sync.WaitGroup

	var {{.ServiceNameLower}}s []models.{{.ServiceName}}
	var totalItems int
	var err error

	wg.Add(2)
	go func() {
		defer wg.Done()
		{{.ServiceNameLower}}s, err = s.Query.FindAll()
	}()
	go func() {
		defer wg.Done()
		totalItems, err = s.Query.Count()
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

func (s *{{.ServiceName}}Service) Create{{.ServiceName}}({{.ServiceNameLower}} models.{{.ServiceName}}) (models.{{.ServiceName}}, error) {
	created{{.ServiceName}}, err := s.Query.Create({{.ServiceNameLower}})
	if err != nil {
		return models.{{.ServiceName}}{}, err
	}

	return *created{{.ServiceName}}, nil
}

func (s *{{.ServiceName}}Service) Get{{.ServiceName}}ByID(ID uint) (models.{{.ServiceName}}, error) {
	{{.ServiceNameLower}}, err := s.Query.FindByID(ID)
	if err != nil {
		return models.{{.ServiceName}}{}, err
	}

	return *{{.ServiceNameLower}}, nil
}

func (s *{{.ServiceName}}Service) Update{{.ServiceName}}(ID uint, {{.ServiceNameLower}}Params models.{{.ServiceName}}) (models.{{.ServiceName}}, error) {
	{{.ServiceNameLower}}, err := s.Query.FindByID(ID)
	if err != nil {
		return models.{{.ServiceName}}{}, err
	}

	{{.ServiceNameLower}}Params.ID = {{.ServiceNameLower}}.ID

	updated{{.ServiceName}}, err := s.Query.Update({{.ServiceNameLower}}Params)
	if err != nil {
		return models.{{.ServiceName}}{}, err
	}

	return *updated{{.ServiceName}}, nil
}

func (s *{{.ServiceName}}Service) Delete{{.ServiceName}}(ID uint) error {
	err := s.Query.DeleteByID(ID)
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
