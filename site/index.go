package site

import (
	"http"
)

type Index struct {
	Filename string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	context := &Index{ Filename:"template/index.html" }
	TemplateRender(w, context.Filename, context)
}

