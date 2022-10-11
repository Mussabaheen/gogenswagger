package template

import (
	"testing"

	"github.com/Mussabaheen/gogenswagger/pkg/swagger"
)

func TestGenerator_GenerateTestFiles(t *testing.T) {

	template := NewGenerator("../../internals/golang.tmpl", "../../generated")
	json := swagger.SwaggerJson{
		Paths: map[string]swagger.Path{
			"get": {
				Get: &swagger.RestApi{
					Description: "This is a comment",
					Tags:        []string{"server"},
					OperationID: "createserverinstance",
					Responses: map[string]swagger.Response{
						"200": {
							Description: "success",
							Schema: swagger.Schema{
								Ref: "successfull",
							},
						},
					},
				},
			},
		},
	}

	template.GenerateTestFiles(&json)
}
