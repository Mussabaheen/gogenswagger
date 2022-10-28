// Package language selects the language
package language

import (
	"log"
)

var (
	nodejs = "templates/nodejs.tmpl"
	golang = "templates/golang.tmpl"
)

// Language handles the selection for language
type Language struct {
	extension string
}

// NewLangugae creates a new Language Service
func NewLangugae(extension string) *Language {
	return &Language{
		extension: extension,
	}
}

// Select returns the language using the extension by user
func (L *Language) Select() (string, string) {
	switch L.extension {
	case "js":
		return nodejs, "js"
	case "go":
		return golang, "go"
	default:
		log.Fatalf("no language selected")
	}
	return "", ""
}
