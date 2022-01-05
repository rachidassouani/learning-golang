package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(rw http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	error := parsedTemplate.Execute(rw, nil)
	if error != nil {
		fmt.Println("Error parsing template: ", error)
		return
	}
}
