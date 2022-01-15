package hanldlers

import (
	"net/http"

	"github.com/rachidassouani/go-project/pkg/render"
)

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.tmpl")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "about.page.tmpl")
}
