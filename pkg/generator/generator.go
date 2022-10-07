package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"

	"github.com/Masterminds/sprig"
	"github.com/Mussabaheen/gotestapi/pkg/fetcher"
)

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
			var responses []string
			for key := range swaggerJson.Paths[k].Get.Responses {
				responses = append(responses, key)
			}
			tempTestCase = append(tempTestCase, TestCase{
				Name:      swaggerJson.Paths[k].Get.OperationID,
				Responses: responses,
			})
			apiTest.Apis[swaggerJson.Paths[k].Get.Tags[0]] = Api{
				PackageName: swaggerJson.Paths[k].Get.Tags[0],
				FileName:    swaggerJson.Paths[k].Get.Tags[0] + "_test.go",
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Put != nil {
			tempTestCase := apiTest.Apis[swaggerJson.Paths[k].Put.Tags[0]].TestCases
			var responses []string
			for key := range swaggerJson.Paths[k].Put.Responses {
				responses = append(responses, key)
			}
			tempTestCase = append(tempTestCase, TestCase{
				Name:      swaggerJson.Paths[k].Put.OperationID,
				Responses: responses,
			})
			apiTest.Apis[swaggerJson.Paths[k].Put.Tags[0]] = Api{
				PackageName: swaggerJson.Paths[k].Put.Tags[0],
				FileName:    swaggerJson.Paths[k].Put.Tags[0] + "_test.go",
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Delete != nil {
			tempTestCase := apiTest.Apis[swaggerJson.Paths[k].Delete.Tags[0]].TestCases
			var responses []string
			for key := range swaggerJson.Paths[k].Delete.Responses {
				responses = append(responses, key)
			}
			tempTestCase = append(tempTestCase, TestCase{
				Name:      swaggerJson.Paths[k].Delete.OperationID,
				Responses: responses,
			})
			apiTest.Apis[swaggerJson.Paths[k].Delete.Tags[0]] = Api{
				PackageName: swaggerJson.Paths[k].Delete.Tags[0],
				FileName:    swaggerJson.Paths[k].Delete.Tags[0] + "_test.go",
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Post != nil {
			tempTestCase := apiTest.Apis[swaggerJson.Paths[k].Post.Tags[0]].TestCases
			var responses []string
			for key := range swaggerJson.Paths[k].Post.Responses {
				responses = append(responses, key)
			}
			tempTestCase = append(tempTestCase, TestCase{
				Name:      swaggerJson.Paths[k].Post.OperationID,
				Responses: responses,
			})
			apiTest.Apis[swaggerJson.Paths[k].Post.Tags[0]] = Api{
				PackageName: swaggerJson.Paths[k].Post.Tags[0],
				FileName:    swaggerJson.Paths[k].Post.Tags[0] + "_test.go",
				TestCases:   tempTestCase,
			}
		}
		if swaggerJson.Paths[k].Update != nil {
			tempTestCase := apiTest.Apis[swaggerJson.Paths[k].Update.Tags[0]].TestCases
			var responses []string
			for key := range swaggerJson.Paths[k].Update.Responses {
				responses = append(responses, key)
			}
			tempTestCase = append(tempTestCase, TestCase{
				Name:      swaggerJson.Paths[k].Update.OperationID,
				Responses: responses,
			})
			apiTest.Apis[swaggerJson.Paths[k].Update.Tags[0]] = Api{
				PackageName: swaggerJson.Paths[k].Update.Tags[0],
				FileName:    swaggerJson.Paths[k].Update.Tags[0] + "_test.go",
				TestCases:   tempTestCase,
			}
		}
	}

	for _, Api := range apiTest.Apis {
		template := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles("../../templates/httpTest.tmpl"))
		var tmplprocessed bytes.Buffer
		err := template.ExecuteTemplate(&tmplprocessed, Api.FileName, Api)
		if err != nil {
			log.Fatalf("Unable to parse data into template: %v\n", err)
		}
		tmplFormatted, err := format.Source(tmplprocessed.Bytes())
		if err != nil {
			log.Fatalf("Could not format processed template: %v\n", err)
		}
		testoutputPath := Api.FileName
		fmt.Println("Writing file: ", testoutputPath)
		f, err := os.Create(testoutputPath)
		if err != nil {
			log.Fatalf("could not save file %v\n", err)
		}
		w := bufio.NewWriter(f)
		w.WriteString(string(tmplFormatted))
		w.Flush()
	}
}
