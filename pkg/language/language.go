// Package language selects the language
package language

import "log"

var (
	nodejs = "internals/templates/nodejs.tmpl"
	golang = "internals/templates/golang.tmpl"
)

// Language handles the selection for language
type Language struct {
	option string
}

// NewLangugae creates a new Language Service
func NewLangugae(option string) *Language {
	return &Language{
		option: option,
	}
}

// Select returns the language using the option by user
func (L *Language) Select() (string, string) {
	switch L.option {
	case "1":
		return nodejs, "js"
	case "2":
		return golang, "go"
	default:
		log.Fatalf("no language selected")
	}
	return "", ""
}
