// package main is the entry point for the gogenswagger
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Mussabaheen/gogenswagger/internals/template"
	"github.com/Mussabaheen/gogenswagger/pkg/language"
	"github.com/Mussabaheen/gogenswagger/pkg/swagger"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("invalid number of arguments, please provide template path and swagger json path")
	}

	jsonFile := os.Args[1]
	if jsonFile == "" {
		log.Fatalf("Swagger json not provided, please provide Swagger JSON file")
	}

	fmt.Println("Kindly select language in which test should be generated. \n 1 : Node.js \n 2 : GoLang ")
	var selectLanguage string
	fmt.Print("Enter your option: ")
	fmt.Scanln(&selectLanguage)

	language := language.NewLangugae(selectLanguage)
	tmplPath, fileExtension := language.Select()

	fetch := swagger.NewSwagger(jsonFile)
	jsonSwagger := fetch.JSONParser()

	testGenerator := template.NewTemplate(tmplPath, "./generated")
	testGenerator.GenerateTestFiles(jsonSwagger, fileExtension)
}
