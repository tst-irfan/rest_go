package lib

import (
	"errors"
)

type FieldArgs struct {
	Name string
	Type string
}

var dataTypes = map[string]string{
	"string": "string",
	"int":    "int",
	"uint":   "uint",
	"float":  "float",
	"time":   "time.Time",
}

func ValidateDataType(dataType string) error {
	if _, ok := dataTypes[dataType]; !ok {
		return errors.New("Invalid data type")
	}
	return nil
}

func NewFieldArgs(name, dataType string) (FieldArgs, error) {
	err := ValidateDataType(dataType)
	if err != nil {
		return FieldArgs{}, err
	}
	return FieldArgs{
		Name: name,
		Type: dataTypes[dataType],
	}, nil
}
