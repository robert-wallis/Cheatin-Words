package site

import (
	"http"
	"template"
)

func TemplateRender(w http.ResponseWriter, filename string, context interface{}) {
	t, err := template.ParseFile(filename, nil)
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
	if err := t.Execute(w, context); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}
