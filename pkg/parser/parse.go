package parser

import (
	"html/template"
	"log"
)

func Parse(tmpl string) *template.Template {
	parsed, err := template.ParseFiles(tmpl)

	if err != nil {
		log.Println("error parsing the template")
	}

	return parsed
}
