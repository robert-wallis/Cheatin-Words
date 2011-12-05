package site

import (
	"http"
	"template"
)

type Index struct {
	T *template.Template
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// prepare variables
	context := &Index{}

	// prepare template
	t, err := template.ParseFile("template/index.html")
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}
	context.T = t

	// display template
	if err := context.T.Execute(w, context); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}
