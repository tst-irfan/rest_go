package lib

import (
	"os"
)

type FileGenerator struct {
	Name string
	Type string
}

func createFile(filePath, content string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	println(currentDir)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
