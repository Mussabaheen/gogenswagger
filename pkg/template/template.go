// Package template uses the test cases to add them into the chosen template
package template

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/Mussabaheen/gogenswagger/pkg/swagger"
)

type Template struct {
	templatePath string // templatePath represents the path of the template being used
	outputFolder string // outputFolder represents the folder for the generated test cases
}

// NewTemplate creates a new Template service
func NewTemplate(templatePath string, outputDestination string) *Template {
	return &Template{
		templatePath: templatePath,
		outputFolder: outputDestination,
	}
}

// GenerateTestFiles generates the test files using the provided template
func (T *Template) GenerateTestFiles(swaggerJson *swagger.SwaggerJson, testExtension string) {
	tmpExtension := filepath.Ext(T.templatePath)
	if tmpExtension != ".tmpl" {
		log.Fatalf("invalid extension provided for template file, file must be *.tmpl")
	}
	apiTest := GeneratedTest{
		GeneratedTests: make(map[string]Test),
	}
	for k := range swaggerJson.Paths {
		if swaggerJson.Paths[k].Get != nil {
			tempTestCase := apiTest.GeneratedTests[swaggerJson.Paths[k].Get.Tags[0]].TestCases
			for key := range swaggerJson.Paths[k].Get.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name:        "Test" + "Get" + strings.ReplaceAll(swaggerJson.Paths[k].Get.OperationID, "-", "_") + "Returns" + key,
					Description: swaggerJson.Paths[k].Get.Description,
				})

			}

			apiTest.GeneratedTests[swaggerJson.Paths[k].Get.Tags[0]] = Test{
				PackageName: swaggerJson.Paths[k].Get.Tags[0],
				FileName:    swaggerJson.Paths[k].Get.Tags[0] + "_test." + testExtension,
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Put != nil {
			tempTestCase := apiTest.GeneratedTests[swaggerJson.Paths[k].Put.Tags[0]].TestCases
			for key := range swaggerJson.Paths[k].Put.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name:        "Test" + "Put" + strings.ReplaceAll(swaggerJson.Paths[k].Put.OperationID, "-", "_") + "Returns" + key,
					Description: swaggerJson.Paths[k].Put.Description,
				})
			}
			apiTest.GeneratedTests[swaggerJson.Paths[k].Put.Tags[0]] = Test{
				PackageName: swaggerJson.Paths[k].Put.Tags[0],
				FileName:    swaggerJson.Paths[k].Put.Tags[0] + "_test." + testExtension,
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Delete != nil {
			tempTestCase := apiTest.GeneratedTests[swaggerJson.Paths[k].Delete.Tags[0]].TestCases
			for key := range swaggerJson.Paths[k].Delete.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name:        "Test" + "Delete" + strings.ReplaceAll(swaggerJson.Paths[k].Delete.OperationID, "-", "_") + "Returns" + key,
					Description: swaggerJson.Paths[k].Delete.Description,
				})
			}
			apiTest.GeneratedTests[swaggerJson.Paths[k].Delete.Tags[0]] = Test{
				PackageName: swaggerJson.Paths[k].Delete.Tags[0],
				FileName:    swaggerJson.Paths[k].Delete.Tags[0] + "_test." + testExtension,
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Post != nil {
			tempTestCase := apiTest.GeneratedTests[swaggerJson.Paths[k].Post.Tags[0]].TestCases
			for key := range swaggerJson.Paths[k].Post.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name:        "Test" + "Post" + strings.ReplaceAll(swaggerJson.Paths[k].Post.OperationID, "-", "_") + "Returns" + key,
					Description: swaggerJson.Paths[k].Post.Description,
				})
			}
			apiTest.GeneratedTests[swaggerJson.Paths[k].Post.Tags[0]] = Test{
				PackageName: swaggerJson.Paths[k].Post.Tags[0],
				FileName:    swaggerJson.Paths[k].Post.Tags[0] + "_test." + testExtension,
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Update != nil {
			tempTestCase := apiTest.GeneratedTests[swaggerJson.Paths[k].Update.Tags[0]].TestCases
			for key := range swaggerJson.Paths[k].Update.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name:        "Test" + "Update" + strings.ReplaceAll(swaggerJson.Paths[k].Update.OperationID, "-", "_") + "Returns" + key,
					Description: swaggerJson.Paths[k].Update.Description,
				})
			}

			apiTest.GeneratedTests[swaggerJson.Paths[k].Update.Tags[0]] = Test{
				PackageName: swaggerJson.Paths[k].Update.Tags[0],
				FileName:    swaggerJson.Paths[k].Update.Tags[0] + "_test." + testExtension,
				TestCases:   tempTestCase,
			}
		}
	}

	for _, Api := range apiTest.GeneratedTests {
		fileName := filepath.Base(T.templatePath)
		tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(T.templatePath))
		var processed bytes.Buffer
		err := tmpl.ExecuteTemplate(&processed, fileName, Api)
		if err != nil {
			log.Fatalf("Unable to parse data into template: %v\n", err)
		}
		var formatted []byte
		if testExtension == "go" {
			formatted, err = format.Source(processed.Bytes())
			if err != nil {
				log.Fatalf("Could not format processed template: %v\n", err)
			}
		} else {
			formatted = processed.Bytes()
		}

		directoryPath := T.outputFolder + "/" + Api.PackageName + "/"
		if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
			log.Fatalf("Could not create directories : %v \n", err)
		}
		outputPath := directoryPath + Api.FileName
		fmt.Println("Writing file: ", outputPath)
		f, err := os.Create(outputPath)
		if err != nil {
			log.Fatalf("could not save file %v\n", err)
		}
		w := bufio.NewWriter(f)
		_, err = w.WriteString(string(formatted))
		if err != nil {
			log.Fatalf("could not write file %v\n", err)
		}
		err = w.Flush()
		if err != nil {
			log.Fatalf("could not flush file %v\n", err)
		}
	}
}
