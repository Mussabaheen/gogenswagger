package template

import "embed"

//go:embed templates/*.tmpl
var Templates embed.FS
