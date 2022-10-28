// package main is the entry point for the gogenswagger
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Mussabaheen/gogenswagger/internals/template"
	"github.com/Mussabaheen/gogenswagger/pkg/language"
	"github.com/Mussabaheen/gogenswagger/pkg/swagger"
)

var (
	fileExtensionUsage = "Specify the language extension, currently supported languages js and go"
	swaggerFileUsage   = "Specify the Swagger JSON file path"
	outputPathUsage    = "Specify the path for generated test packages"
)

func main() {
	// fileExtension represents the arg with flag -l
	fileExtension := flag.String("l", "go", fileExtensionUsage)

	// swaggerFile represents the arg with flag -s
	swaggerFile := flag.String("s", "", swaggerFileUsage)

	// outputPath represents the arg with flat -o
	outputPath := flag.String("o", "./generated", outputPathUsage)
	flag.Parse()
	fmt.Println(*fileExtension, *swaggerFile)

	if *swaggerFile == "" {
		log.Fatalf("Swagger json not provided, please provide Swagger JSON file with -s flag")
	}

	var selectLanguage string

	language := language.NewLangugae(selectLanguage)
	tmplPath, fileExtn := language.Select()

	fetch := swagger.NewSwagger(*swaggerFile)
	jsonSwagger := fetch.JSONParser()

	testGenerator := template.NewTemplate(tmplPath, *outputPath)
	testGenerator.GenerateTestFiles(jsonSwagger, fileExtn)
}
