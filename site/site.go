package site

import (
	"fmt"
	"http"
	"scrabble"
)

func init() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	if 0 != len(query) {
		strings := scrabble.StringPermutations(query)
		for s := range strings {
			fmt.Fprint(w, "%s<br>\n", s)
		}
	}
	renderTemplate(w, "template/index.html")
	fmt.Fprint(w, "site page")
}

func renderTemplate(w http.ResponseWriter, filename string) {
	page := Page{}
	t, err := page.Load(filename)
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
	if err := t.Execute(w, page); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}

