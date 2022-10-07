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

func (G *Generator) GenerateTestFiles(swaggerJson fetcher.SwaggerJson) {
	apiTest := ApiTest{
		Apis: []Api{},
	}
	for _, Api := range apiTest.Apis {
		template := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles("./templates/httpTest.tmpl"))
		var tmplprocessed bytes.Buffer
		err := template.ExecuteTemplate(&tmplprocessed, Api.FileName, Api)
		if err != nil {
			log.Fatalf("Unable to parse data into template: %v\n", err)
		}
		tmplFormatted, err := format.Source(tmplprocessed.Bytes())
		if err != nil {
			log.Fatalf("Could not format processed template: %v\n", err)
		}
		testoutputPath := Api.FileName + "_test.go"
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
