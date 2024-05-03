package main

import (
	"fmt"
	"log"
	"os"
	"rest_go/generators/lib"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide type and name")
		return
	}

	typeName := args[1]
	name := args[2]
	fieldArgs := []lib.FieldArgs{}
	

	if len(args) > 2 {
		for i := 3; i < len(args); i++ {
			arg := args[i]
			argSplit :=  strings.Split(arg, ":")
			field, err := lib.NewFieldArgs(argSplit[0], argSplit[1])
			if err != nil {
				log.Fatal(err)
			}
			fieldArgs = append(fieldArgs, field)
		}
	}

	fileGenerator := lib.FileGenerator{
		Name: name,
		Type: typeName,
	}

	switch typeName {
	case "model":
		err := lib.GenerateModel(fileGenerator, fieldArgs)
		if err != nil {
			log.Fatal(err)
		}
	case "service":
		err := lib.GenerateService(fileGenerator)
		if err != nil {
			log.Fatal(err)
		}
	case "controller":
		err := lib.GenerateController(fileGenerator)
		if err != nil {
			log.Fatal(err)
		}
	case "router":
		err := lib.GenerateRouter(fileGenerator)
		if err != nil {
			log.Fatal(err)
		}
	case "scaffold":
		err := lib.GenerateModel(fileGenerator, fieldArgs)
		err =  lib.GenerateService(fileGenerator)
		err =  lib.GenerateController(fileGenerator)
		err =  lib.GenerateRouter(fileGenerator)
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Unknown type")
	}

}
