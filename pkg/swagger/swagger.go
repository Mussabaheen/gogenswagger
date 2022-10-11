// Package swagger reads the swagger json file.
package swagger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
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

// JsonParser loads and Parses the JSON file into SwaggerJson struct
func (S *Swagger) JsonParser() *SwaggerJson {
	jsonFile, err := os.Open(S.fileLocation)
	if err != nil {
		log.Fatalf("unable to open swagger JSON file %v\n", err)
	}
	fmt.Println("file loaded succesfully!")
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
