package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	filePath := "./docs/docs.go"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("File does not exist: %s\n", filePath)
		return
	}

	fmt.Printf("Processing file: %s\n", filePath)
	err := replaceInFile(filePath)
	if err != nil {
		fmt.Printf("Error replacing content: %v\n", err)
	}
}

func replaceInFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	newContent := strings.ReplaceAll(string(content),
		"swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)",
		"swag.Register(SwaggerInfo.InstanceName(), &SwaggerInfoWrapper{*SwaggerInfo})")

	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Replaced content in file: %s\n", filePath)
	return nil
}
