package handler

import (
	"net/http"

	"github.com/IsroilMukhitdinov/go_project/pkg/entities"
	"github.com/IsroilMukhitdinov/go_project/pkg/parser"
)

func Handle(response http.ResponseWriter, pageData entities.PageData) {

	parsed := parser.Parse("./templates/page.tmpl")
	parsed.Execute(response, pageData)
}
