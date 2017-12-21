package main

import (
	"os"
	"fmt"
)

func getSchema(schema string) (string, error) {
	pwd, _ := os.Getwd()
	schemaFile := fmt.Sprintf("%s/schema/%s.json", pwd, schema)

	if _, err := os.Stat(schemaFile); os.IsNotExist(err) {
		return "", err
	}

	return schemaFile, nil
}
