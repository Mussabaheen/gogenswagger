package language

import "log"

var (
	nodejs = "internals/templates/nodejs.tmpl"
	golang = "internals/templates/golang.tmpl"
)

type Language struct {
	option string
}

func NewLangugae(option string) *Language {
	return &Language{
		option: option,
	}
}

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
