package lib

import (
	"log"
	"os"
	"strings"
)

type FileGenerator struct {
	Name string
	Type string
}

func createFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	log.Printf("File %s created successfully", filePath)
	return nil
}

func AppendContent(filePath, content string, fileType string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	updatedFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer updatedFile.Close()

	currentContent := string(data)
	var toReplace string
	if fileType == "model" {
		toReplace = `
	)
}`
	} else if fileType == "router" {
		toReplace = `
	return r
}`
	}
	newContent := strings.Replace(currentContent, toReplace, content+toReplace, -1)
	_, err = updatedFile.WriteString(newContent)
	if err != nil {
		return err
	}
	log.Printf("File %s updated successfully", filePath)

	return nil
}
