package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/Mussabaheen/gotestapi/pkg/fetcher"
)

type Data struct {
	PackageName  string
	FunctionName string
}
type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (G *Generator) GenerateTestFiles(swaggerJson *fetcher.SwaggerJson) {
	apiTest := ApiTest{
		Apis: make(map[string]Api),
	}
	for k := range swaggerJson.Paths {
		if swaggerJson.Paths[k].Get != nil {
			tempTestCase := apiTest.Apis[swaggerJson.Paths[k].Get.Tags[0]].TestCases
			for key := range swaggerJson.Paths[k].Get.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name: "Test" + strings.ReplaceAll(swaggerJson.Paths[k].Get.OperationID, "-", "_") + "Returns" + key})
			}

			apiTest.Apis[swaggerJson.Paths[k].Get.Tags[0]] = Api{
				PackageName: swaggerJson.Paths[k].Get.Tags[0],
				FileName:    swaggerJson.Paths[k].Get.Tags[0] + "_test.go",
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Put != nil {
			tempTestCase := apiTest.Apis[swaggerJson.Paths[k].Put.Tags[0]].TestCases
			for key := range swaggerJson.Paths[k].Put.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name: "Test" + strings.ReplaceAll(swaggerJson.Paths[k].Put.OperationID, "-", "_") + "Returns" + key})
			}
			apiTest.Apis[swaggerJson.Paths[k].Put.Tags[0]] = Api{
				PackageName: swaggerJson.Paths[k].Put.Tags[0],
				FileName:    swaggerJson.Paths[k].Put.Tags[0] + "_test.go",
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Delete != nil {
			tempTestCase := apiTest.Apis[swaggerJson.Paths[k].Delete.Tags[0]].TestCases
			for key := range swaggerJson.Paths[k].Delete.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name: "Test" + strings.ReplaceAll(swaggerJson.Paths[k].Delete.OperationID, "-", "_") + "Returns" + key})
			}
			apiTest.Apis[swaggerJson.Paths[k].Delete.Tags[0]] = Api{
				PackageName: swaggerJson.Paths[k].Delete.Tags[0],
				FileName:    swaggerJson.Paths[k].Delete.Tags[0] + "_test.go",
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Post != nil {
			tempTestCase := apiTest.Apis[swaggerJson.Paths[k].Post.Tags[0]].TestCases
			for key := range swaggerJson.Paths[k].Post.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name: "Test" + strings.ReplaceAll(swaggerJson.Paths[k].Post.OperationID, "-", "_") + "Returns" + key})
			}
			apiTest.Apis[swaggerJson.Paths[k].Post.Tags[0]] = Api{
				PackageName: swaggerJson.Paths[k].Post.Tags[0],
				FileName:    swaggerJson.Paths[k].Post.Tags[0] + "_test.go",
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Update != nil {
			tempTestCase := apiTest.Apis[swaggerJson.Paths[k].Update.Tags[0]].TestCases
			for key := range swaggerJson.Paths[k].Update.Responses {
				tempTestCase = append(tempTestCase, TestCase{
					Name: "Test" + strings.ReplaceAll(swaggerJson.Paths[k].Update.OperationID, "-", "_") + "Returns" + key})
			}

			apiTest.Apis[swaggerJson.Paths[k].Update.Tags[0]] = Api{
				PackageName: swaggerJson.Paths[k].Update.Tags[0],
				FileName:    swaggerJson.Paths[k].Update.Tags[0] + "_test.go",
				TestCases:   tempTestCase,
			}
		}
	}

	for _, Api := range apiTest.Apis {
		fileName := "templates/httpTest.tmpl"
		tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(fileName))
		var processed bytes.Buffer
		err := tmpl.ExecuteTemplate(&processed, "httpTest.tmpl", Api)
		if err != nil {
			log.Fatalf("Unable to parse data into template: %v\n", err)
		}
		formatted, err := format.Source(processed.Bytes())
		if err != nil {
			log.Fatalf("Could not format processed template: %v\n", err)
		}
		directoryPath := "generated/" + Api.PackageName + "/"
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
		w.WriteString(string(formatted))
		w.Flush()
	}
}
