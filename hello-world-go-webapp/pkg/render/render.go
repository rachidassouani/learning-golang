package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(rw http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	error := parsedTemplate.Execute(rw, nil)
	if error != nil {
		fmt.Println("Error parsing template: ", error)
		return
	}
}
