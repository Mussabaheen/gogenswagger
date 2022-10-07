package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"

	"github.com/Masterminds/sprig"
)

type Data struct {
	PackageName  string
	FunctionName string
}

func main() {
	data := Data{
		PackageName:  "main",
		FunctionName: "Testletsdosomefuckingtests",
	}
	processTemplate("iBuilder.tmpl", "iBuilder_test.go", data)

}
func processTemplate(fileName string, outputFile string, data Data) {
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(fileName))
	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, fileName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}
	outputPath := outputFile
	fmt.Println("Writing file: ", outputPath)
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("could not save file %v\n", err)
	}
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()
}
