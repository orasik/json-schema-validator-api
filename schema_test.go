package main

import (
	"testing"
	"os"
	"fmt"
)

func TestValidGetSchema(t *testing.T) {

	schemaFile, err := getSchema("person")

	if err != nil {
		t.Errorf("Found error for a valid schema! %s", err)
	}

	pwd, _ := os.Getwd()
	expectedFile := fmt.Sprintf("%s/schema/person.json", pwd)

	if expectedFile != schemaFile {
		t.Errorf("Expected file %s, while result is: %s", expectedFile, schemaFile)
	}

}

func TestMissingSchemaFile(t *testing.T) {

	schemaFile, err := getSchema("blabla")

	if err == nil {
		t.Error("Error expected, got nil")
	}

	expectedFile := ""

	if expectedFile != schemaFile {
		t.Errorf("Expected file %s, while result is: %s", expectedFile, schemaFile)
	}

}
