package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Mussabaheen/gogenswagger/pkg/language"
	"github.com/Mussabaheen/gogenswagger/pkg/swagger"
	"github.com/Mussabaheen/gogenswagger/pkg/template"
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
	templatePath, fileExtension := language.Select()

	fetch := swagger.NewSwagger(jsonFile)
	jsonSwagger := fetch.JsonParser()

	testGenerator := template.NewTemplate(templatePath, "./generated")
	testGenerator.GenerateTestFiles(jsonSwagger, fileExtension)
}
