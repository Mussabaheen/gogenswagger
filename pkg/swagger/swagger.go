// Package swagger reads the swagger json file.
package swagger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Swagger struct {
	fileLocation string
}

// NewSwagger creates a new Swagger service
func NewSwagger(fileLocation string) *Swagger {
	return &Swagger{
		fileLocation: fileLocation,
	}
}

// JsonParser checks, loads and Parses the JSON file into SwaggerJson struct
func (S *Swagger) JsonParser() *SwaggerJson {
	fileExtension := filepath.Ext(S.fileLocation)
	if fileExtension != ".json" {
		log.Fatalf("invalid extension provided for swagger file, file must be *.json")
	}

	jsonFile, err := os.Open(S.fileLocation)
	if err != nil {
		log.Fatalf("unable to open swagger JSON file %v\n", err)
	}
	fmt.Println("JSON file loaded succesfully!")
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("unable to read data from swagger JSON file %v\n", err)
	}
	var swagger SwaggerJson

	err = json.Unmarshal(byteValue, &swagger)

	if err != nil {
		log.Fatalf("unable to load data into SwaggerJson struct %v\n", err)
	}
	return &swagger
}
