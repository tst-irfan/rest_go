package lib

import (
	"fmt"
	"path/filepath"
	"strings"
)

const serviceDir = "app/services"

func GenerateService(input FileGenerator) error {
	servicePath := filepath.Join(serviceDir, fmt.Sprintf("%s.service.go", strings.ToLower(input.Name)))
	serviceName := strings.Title(input.Name)
	content := ServiceTemplate(serviceName, strings.ToLower(input.Name))
	err := createFile(servicePath, content)
	if err != nil {
		return err
	}
	return nil
}

func ServiceTemplate(serviceName string, serviceNameLower string) string {
	content := `package services

	import (
		"rest_go/app/models"
		"rest_go/app/types"
	)
	
	func ShowAll%S() ([]models.%S, error, types.MetaData) {
		%ss, err := models.%SQuery.FindAll()
		if err != nil {
			return nil, err, types.MetaData{}
		}
		totalItems, err := models.%SQuery.Count()
		if err != nil {
			return nil, err, types.MetaData{}
		}
		metaData := types.MetaData{
			TotalItems: totalItems,
		}
	
		return %ss, nil, metaData
	}
	
	func Create%S(%s models.%S) (models.%S, error) {
		created%S, err := models.%SQuery.Create(%s)
		if err != nil {
			return models.%S{}, err
		}
	
		return *created%S, nil
	}
	
	func Get%SByID(ID uint) (models.%S, error) {
		%s, err := models.%SQuery.FindByID(ID)
		if err != nil {
			return models.%S{}, err
		}
	
		return *%s, nil
	}
	
	func Update%S(ID uint, %sParams models.%S) (models.%S, error) {
		%s, err := models.%SQuery.FindByID(ID)
		if err != nil {
			return models.%S{}, err
		}
	
		%sParams.ID = %s.ID
	
		updated%S, err := models.%SQuery.Update(%sParams)
		if err != nil {
			return models.%S{}, err
		}
	
		return *updated%S, nil
	}
	
	func Delete%S(ID uint) error {
		err := models.%SQuery.DeleteByID(ID)
		if err != nil {
			return err
		}
	
		return nil
	}`

	content = strings.Replace(content, "%S", serviceName, -1)
	content = strings.Replace(content, "%s", serviceNameLower, -1)
	return content
}
