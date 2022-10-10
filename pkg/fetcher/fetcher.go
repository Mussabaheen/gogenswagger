package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Fetcher struct {
	fileLocation string
}

func NewFetcher(fileLocation string) *Fetcher {
	return &Fetcher{
		fileLocation: fileLocation,
	}
}

func (f *Fetcher) JsonParser() *SwaggerJson {
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
