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

// Template handles the generation of test cases
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
func (T *Template) GenerateTestFiles(swaggerJSON *swagger.JSON, testExtension string) {
	tmpExtension := filepath.Ext(T.templatePath)
	if tmpExtension != ".tmpl" {
		log.Fatalf("invalid extension provided for template file, file must be *.tmpl")
	}
	apiTest := GeneratedTest{
		GeneratedTests: make(map[string]Test),
	}
	packageName := "test"
	for k := range swaggerJSON.Paths {
		if swaggerJSON.Paths[k].Get != nil {
			if len(swaggerJSON.Paths[k].Get.Tags) != 0 {
				packageName = swaggerJSON.Paths[k].Get.Tags[0]
			}
			tempTestCase := apiTest.GeneratedTests[packageName].TestCases
			for key := range swaggerJSON.Paths[k].Get.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name:         "Test" + "Get" + strings.ReplaceAll(swaggerJSON.Paths[k].Get.OperationID, "-", "_") + "Returns" + key,
					Description:  swaggerJSON.Paths[k].Get.Description,
					Endpoint:     k,
					ResponseCode: key,
					Method:       "Get",
				})

			}

			apiTest.GeneratedTests[packageName] = Test{
				PackageName: packageName,
				FileName:    packageName + "_test." + testExtension,
				TestCases:   tempTestCase,
			}
		}
		if swaggerJSON.Paths[k].Put != nil {
			if len(swaggerJSON.Paths[k].Put.Tags) != 0 {
				packageName = swaggerJSON.Paths[k].Put.Tags[0]
			}
			tempTestCase := apiTest.GeneratedTests[packageName].TestCases
			for key := range swaggerJSON.Paths[k].Put.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name:         "Test" + "Put" + strings.ReplaceAll(swaggerJSON.Paths[k].Put.OperationID, "-", "_") + "Returns" + key,
					Description:  swaggerJSON.Paths[k].Put.Description,
					Endpoint:     k,
					ResponseCode: key,
					Method:       "Put",
				})
			}
			apiTest.GeneratedTests[packageName] = Test{
				PackageName: packageName,
				FileName:    packageName + "_test." + testExtension,
				TestCases:   tempTestCase,
			}
		}
		if swaggerJSON.Paths[k].Delete != nil {
			if len(swaggerJSON.Paths[k].Delete.Tags) != 0 {
				packageName = swaggerJSON.Paths[k].Delete.Tags[0]
			}
			tempTestCase := apiTest.GeneratedTests[packageName].TestCases
			for key := range swaggerJSON.Paths[k].Delete.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name:         "Test" + "Delete" + strings.ReplaceAll(swaggerJSON.Paths[k].Delete.OperationID, "-", "_") + "Returns" + key,
					Description:  swaggerJSON.Paths[k].Delete.Description,
					Endpoint:     k,
					ResponseCode: key,
					Method:       "Delete",
				})
			}
			apiTest.GeneratedTests[packageName] = Test{
				PackageName: packageName,
				FileName:    packageName + "_test." + testExtension,
				TestCases:   tempTestCase,
			}
		}
		if swaggerJSON.Paths[k].Post != nil {
			if len(swaggerJSON.Paths[k].Post.Tags) != 0 {
				packageName = swaggerJSON.Paths[k].Post.Tags[0]
			}
			tempTestCase := apiTest.GeneratedTests[packageName].TestCases
			for key := range swaggerJSON.Paths[k].Post.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name:         "Test" + "Post" + strings.ReplaceAll(swaggerJSON.Paths[k].Post.OperationID, "-", "_") + "Returns" + key,
					Description:  swaggerJSON.Paths[k].Post.Description,
					Endpoint:     k,
					ResponseCode: key,
					Method:       "Post",
				})
			}
			apiTest.GeneratedTests[packageName] = Test{
				PackageName: packageName,
				FileName:    packageName + "_test." + testExtension,
				TestCases:   tempTestCase,
			}
		}
		if swaggerJSON.Paths[k].Update != nil {
			tempTestCase := apiTest.GeneratedTests[swaggerJSON.Paths[k].Update.Tags[0]].TestCases
			for key := range swaggerJSON.Paths[k].Update.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name:         "Test" + "Update" + strings.ReplaceAll(swaggerJSON.Paths[k].Update.OperationID, "-", "_") + "Returns" + key,
					Description:  swaggerJSON.Paths[k].Update.Description,
					Endpoint:     k,
					ResponseCode: key,
					Method:       "Update",
				})
			}

			apiTest.GeneratedTests[swaggerJSON.Paths[k].Update.Tags[0]] = Test{
				PackageName: swaggerJSON.Paths[k].Update.Tags[0],
				FileName:    swaggerJSON.Paths[k].Update.Tags[0] + "_test." + testExtension,
				TestCases:   tempTestCase,
			}
		}
	}

	for _, API := range apiTest.GeneratedTests {
		fileName := filepath.Base(T.templatePath)
		tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(T.templatePath))
		var processed bytes.Buffer
		err := tmpl.ExecuteTemplate(&processed, fileName, API)
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

		directoryPath := T.outputFolder + "/" + API.PackageName + "/"
		if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
			log.Fatalf("Could not create directories : %v \n", err)
		}
		outputPath := directoryPath + API.FileName
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
