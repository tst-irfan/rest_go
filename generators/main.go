package main

import (
	"fmt"
	"log"
	"os"
	"rest_go/generators/lib"
	"rest_go/generators/types"
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
	fieldArgs := []types.FieldArgs{}

	if len(args) > 2 {
		for i := 3; i < len(args); i++ {
			arg := args[i]
			argSplit := strings.Split(arg, ":")
			field, err := types.NewFieldArgs(argSplit[0], argSplit[1])
			if err != nil {
				log.Fatal(err)
			}
			fieldArgs = append(fieldArgs, field)
		}
	}

	switch typeName {
	case "model":
		err := lib.GenerateModel(name, fieldArgs)
		if err != nil {
			log.Fatal(err)
		}
	case "service":
		err := lib.GenerateService(name)
		if err != nil {
			log.Fatal(err)
		}
	case "controller":
		err := lib.GenerateController(name)
		if err != nil {
			log.Fatal(err)
		}
	case "router":
		err := lib.GenerateRouter(name)
		if err != nil {
			log.Fatal(err)
		}
	case "scaffold":
		err := lib.GenerateModel(name, fieldArgs)
		err = lib.GenerateService(name)
		err = lib.GenerateController(name)
		err = lib.GenerateRouter(name)
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Unknown type")
	}

}
