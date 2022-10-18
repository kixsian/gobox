package handlers

import (
	"net/http"

	"github.com/kixsian/gobox/pkg/render"
)

func Root(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "root.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.page.html")
}
