package main

import (
	"fmt"
	"log"
	"os"
	"rest_go/generators/lib"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide type and name")
		return
	}

	typeName := os.Args[1]
	name := os.Args[2]

	fileGenerator := lib.FileGenerator{
		Name: name,
		Type: typeName,
	}

	switch typeName {
	case "model":
		err := lib.GenerateModel(fileGenerator)
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
		err := lib.GenerateModel(fileGenerator)
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
