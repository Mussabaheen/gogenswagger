package swagger

import (
	"fmt"

	"github.com/Mussabaheen/gotestapi/pkg/fetcher"
)

type Swagger struct {
	docs fetcher.SwaggerJson
}

func NewSwagger(docs fetcher.SwaggerJson) *Swagger {
	return &Swagger{
		docs: docs,
	}
}

func (S *Swagger) ParseSwaggerJson() {
	fmt.Println(S.docs.Paths)
}
