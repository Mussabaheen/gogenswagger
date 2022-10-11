package main

import (
	"log"
	"os"

	"github.com/Mussabaheen/gogenswagger/pkg/swagger"
	"github.com/Mussabaheen/gogenswagger/pkg/template"
)

func main() {

	templatePath := os.Args[1]
	if templatePath == "" {
		log.Fatalf("template path cannot be empty")
	}

	jsonFile := os.Args[2]
	if jsonFile == "" {
		log.Fatalf("Swagger json not provided, please provide Swagger JSON file")
	}

	fetch := swagger.NewSwagger(jsonFile)
	jsonSwagger := fetch.JsonParser()
	testGenerator := template.NewTemplate(templatePath, "./generated")
	testGenerator.GenerateTestFiles(jsonSwagger)
}
