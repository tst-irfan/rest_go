package helpers

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Field struct {
	Name  string
	Value string
	Type  string
}

type ValidationHelper struct {
	RequiredFields          []string
	ShouldGreaterThanFields []Field
	ShouldLessThanFields    []Field
}

func (v *ValidationHelper) Validate(T interface{}) (bool, error) {
	data := structToFields(T)

	err := validateRequiredFields(data, v.RequiredFields)
	if err != nil {
		return false, err
	}
	err = validateShouldGreaterThanFields(data, v.ShouldGreaterThanFields)
	if err != nil {
		return false, err
	}
	err = validateShouldLessThanFields(data, v.ShouldLessThanFields)
	if err != nil {
		return false, err
	}

	return true, nil
}

func validateRequiredFields(data []Field, requiredFields []string) error {
	for _, field := range requiredFields {
		found := false
		for _, f := range data {
			if f.Name == field {
				found = f.Value != ""
				break
			}
		}
		if !found {
			return fmt.Errorf("required field '%s' is missing", field)
		}
	}

	return nil
}

func validateShouldGreaterThanFields(data []Field, shouldGreaterThanFields []Field) error {
	for _, field := range shouldGreaterThanFields {
		for _, f := range data {
			if f.Name == field.Name {
				value, _ := strconv.ParseFloat(f.Value, 64)
				comparisonValue, _ := strconv.ParseFloat(field.Value, 64)
				if value <= comparisonValue {
					return fmt.Errorf("value of field '%s' should be greater than %f", f.Name, field.Value)
				}
				break
			}
		}
	}

	return nil
}

func validateShouldLessThanFields(data []Field, shouldLessThanFields []Field) error {
	for _, field := range shouldLessThanFields {
		for _, f := range data {
			if f.Name == field.Name {
				value, _ := strconv.ParseFloat(f.Value, 64)
				comparisonValue, _ := strconv.ParseFloat(field.Value, 64)
				if value >= comparisonValue {
					return fmt.Errorf("value of field '%s' should be less than %f", f.Name, formattedValue(field.Value, field.Type))
				}
				break
			}
		}
	}

	return nil
}

func structToFields(data interface{}) []Field {
	var fields []Field
	s := reflect.ValueOf(data)

	if s.Kind() != reflect.Ptr {
		s = reflect.New(reflect.TypeOf(data))
	}
	s = s.Elem()

	for i := 0; i < s.NumField(); i++ {
		field := s.Type().Field(i)
		value := fmt.Sprintf("%v", s.Field(i).Interface())
		fieldType := field.Type.String()

		if field.Type.Kind() == reflect.Struct && field.Type != reflect.TypeOf(time.Time{}) {
			nestedFields := structToFields(s.Field(i).Interface())
			for _, nf := range nestedFields {
				nf.Name = field.Name + "." + nf.Name
				nf.Type = fieldType
				fields = append(fields, nf)
			}
		} else {
			if field.Type == reflect.TypeOf(time.Time{}) {
				value = fmt.Sprintf("%v", s.Field(i).Interface().(time.Time).Unix())
			}
			fields = append(fields, Field{Name: field.Name, Value: value, Type: fieldType})
		}
	}

	return fields
}

func formattedValue(value string, dataType string) string {
	if dataType == "time" {
		t, _ := strconv.ParseInt(value, 10, 64)
		return time.Unix(t, 0).String()
	}
	return value
}
