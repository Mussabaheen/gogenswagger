package main

import (
	"github.com/Mussabaheen/gotestapi/pkg/fetcher"
	"github.com/Mussabaheen/gotestapi/pkg/generator"
)

func main() {
	fetch := fetcher.NewFetcher("templates/docs.json")
	jsonSwagger := fetch.JsonParser()
	testGenerator := generator.NewGenerator()
	testGenerator.GenerateTestFiles(jsonSwagger)
}
