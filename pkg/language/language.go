// Package language selects the language
package language

import "log"

var (
	nodejs = "templates/nodejs.tmpl"
	golang = "templates/golang.tmpl"
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
	case "js":
		return nodejs, "js"
	case "go":
		return golang, "go"
	default:
		log.Fatalf("no language selected")
	}
	return "", ""
}
