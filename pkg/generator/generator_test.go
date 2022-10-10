package generator

import (
	"testing"

	"github.com/Mussabaheen/gotestapi/pkg/fetcher"
)

func TestGenerator_GenerateTestFiles(t *testing.T) {

	generator := NewGenerator("../../internals/golang.tmpl", "../../generated")
	json := fetcher.SwaggerJson{
		Paths: map[string]fetcher.Path{
			"get": {
				Get: &fetcher.RestApi{
					Description: "This is a comment",
					Tags:        []string{"server"},
					OperationID: "createserverinstance",
					Responses: map[string]fetcher.Response{
						"200": {
							Description: "success",
							Schema: fetcher.Schema{
								Ref: "successfull",
							},
						},
					},
				},
			},
		},
	}

	generator.GenerateTestFiles(&json)
}
