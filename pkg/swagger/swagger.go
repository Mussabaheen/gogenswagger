// Package swagger reads the swagger json file.
package swagger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Swagger struct {
	fileLocation string
}

func NewSwagger(fileLocation string) *Swagger {
	return &Swagger{
		fileLocation: fileLocation,
	}
}

func (f *Swagger) JsonParser() *SwaggerJson {
	jsonFile, err := os.Open(f.fileLocation)
	if err != nil {
		panic(err)
	}
	fmt.Println("file loaded succesfully!")
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	var swagger SwaggerJson

	err = json.Unmarshal(byteValue, &swagger)

	if err != nil {
		panic(err)
	}
	return &swagger
}
