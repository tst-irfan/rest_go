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
		if err := lib.Generate(name, "model", fieldArgs); err != nil {
			log.Fatal(err)
		}
		if err := lib.Generate(name, "request", fieldArgs); err != nil {
			log.Fatal(err)
		}
	case "service":
		if err := lib.Generate(name, "service", nil); err != nil {
			log.Fatal(err)
		}
	case "controller":
		if err := lib.Generate(name, "controller", nil); err != nil {
			log.Fatal(err)
		}
	case "router":
		if err := lib.Generate(name, "router", nil); err != nil {
			log.Fatal(err)
		}
	case "scaffold":
		if err := lib.Generate(name, "model", fieldArgs); err != nil {
			log.Fatal(err)
		}
		if err := lib.Generate(name, "service", nil); err != nil {
			log.Fatal(err)
		}
		if err := lib.Generate(name, "controller", nil); err != nil {
			log.Fatal(err)
		}
		if err := lib.Generate(name, "router", nil); err != nil {
			log.Fatal(err)
		}
		if err := lib.Generate(name, "request", fieldArgs); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Unknown type")
	}
	

}
