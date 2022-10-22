package template

import (
	"io"
	"os"
	"testing"

	"github.com/Mussabaheen/gogenswagger/pkg/swagger"
	"github.com/stretchr/testify/assert"
)

func TestTemplate_GetGenerateTestFiles(t *testing.T) {
	template := NewTemplate("templates/golang.tmpl", "../../generated")
	json := swagger.JSON{
		Paths: map[string]swagger.Path{
			"get": {
				Get: &swagger.RestAPI{
					Description: "This is a comment",
					Tags:        []string{"server"},
					OperationID: "createserverinstance",
					Responses: map[string]swagger.Response{
						"200": {
							Description: "success",
						},
					},
				},
			},
		},
	}
	template.GenerateTestFiles(&json, "go")
	goFile, err := os.Open("../../generated/server/server_test.go")
	assert.Nil(t, err, "unable to open generated file")
	defer goFile.Close()
	defer os.Remove("../../generated/server/server_test.go")
	_, err = io.ReadAll(goFile)
	assert.Nil(t, err, "unable to open file")
}
func TestTemplate_PostGenerateTestFiles(t *testing.T) {
	template := NewTemplate("templates/golang.tmpl", "../../generated")
	json := swagger.JSON{
		Paths: map[string]swagger.Path{
			"post": {
				Post: &swagger.RestAPI{
					Description: "This is a comment",
					Tags:        []string{"server"},
					OperationID: "createserverinstance",
					Responses: map[string]swagger.Response{
						"200": {
							Description: "success",
						},
					},
				},
			},
		},
	}
	template.GenerateTestFiles(&json, "go")
	goFile, err := os.Open("../../generated/server/server_test.go")
	assert.Nil(t, err, "unable to open generated file")
	defer goFile.Close()
	defer os.Remove("../../generated/server/server_test.go")
	_, err = io.ReadAll(goFile)
	assert.Nil(t, err, "unable to open file")
}

func TestTemplate_UpdateGenerateTestFiles(t *testing.T) {
	template := NewTemplate("templates/golang.tmpl", "../../generated")
	json := swagger.JSON{
		Paths: map[string]swagger.Path{
			"update": {
				Update: &swagger.RestAPI{
					Description: "This is a comment",
					Tags:        []string{"server"},
					OperationID: "createserverinstance",
					Responses: map[string]swagger.Response{
						"200": {
							Description: "success",
						},
					},
				},
			},
		},
	}
	template.GenerateTestFiles(&json, "go")
	goFile, err := os.Open("../../generated/server/server_test.go")
	assert.Nil(t, err, "unable to open generated file")
	defer goFile.Close()
	defer os.Remove("../../generated/server/server_test.go")
	_, err = io.ReadAll(goFile)
	assert.Nil(t, err, "unable to open file")
}
func TestTemplate_DeleteGenerateTestFiles(t *testing.T) {
	template := NewTemplate("templates/golang.tmpl", "../../generated")
	json := swagger.JSON{
		Paths: map[string]swagger.Path{
			"Delete": {
				Delete: &swagger.RestAPI{
					Description: "This is a comment",
					Tags:        []string{"server"},
					OperationID: "createserverinstance",
					Responses: map[string]swagger.Response{
						"200": {
							Description: "success",
						},
					},
				},
			},
		},
	}
	template.GenerateTestFiles(&json, "go")
	goFile, err := os.Open("../../generated/server/server_test.go")
	assert.Nil(t, err, "unable to open generated file")
	defer goFile.Close()
	defer os.Remove("../../generated/server/server_test.go")
	_, err = io.ReadAll(goFile)
	assert.Nil(t, err, "unable to open file")
}
func TestTemplate_PutGenerateTestFiles(t *testing.T) {
	template := NewTemplate("templates/golang.tmpl", "../../generated")
	json := swagger.JSON{
		Paths: map[string]swagger.Path{
			"put": {
				Put: &swagger.RestAPI{
					Description: "This is a comment",
					Tags:        []string{"server"},
					OperationID: "createserverinstance",
					Responses: map[string]swagger.Response{
						"200": {
							Description: "success",
						},
					},
				},
			},
		},
	}
	template.GenerateTestFiles(&json, "go")
	goFile, err := os.Open("../../generated/server/server_test.go")
	assert.Nil(t, err, "unable to open generated file")
	defer goFile.Close()
	defer os.Remove("../../generated/server/server_test.go")
	_, err = io.ReadAll(goFile)
	assert.Nil(t, err, "unable to open file")
}
