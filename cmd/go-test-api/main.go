package main

import (
	"log"
	"os"

	"github.com/Mussabaheen/gotestapi/pkg/fetcher"
	"github.com/Mussabaheen/gotestapi/pkg/generator"
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

	fetch := fetcher.NewFetcher(jsonFile)
	jsonSwagger := fetch.JsonParser()
	testGenerator := generator.NewGenerator(templatePath, "./generated")
	testGenerator.GenerateTestFiles(jsonSwagger)
}
