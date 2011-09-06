package site

import (
	"fmt"
	"http"
	"strings"
	"word"
)

type Search struct {
	Filename string
	Q string
	Permutations []string
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")

	// max 8 letters
	if len(query) > 8 {
		query = query[0:8]
		http.Redirect(w, r, fmt.Sprintf("search?q=%s", query), 302)
		return
	}
	// lowercase only
	if query != strings.ToLower(query) {
		http.Redirect(w, r, fmt.Sprintf("search?q=%s", strings.ToLower(query)), 302)
		return
	}

	context := &Search{
		Filename: "template/search.html",
		Q: query,
	}
	if 0 != len(query) {
		if len(query) > 8 {
			query = query[0:8]
		}
		channel := word.StringPermutations(query)
		for p := range channel {
			context.Permutations = append(context.Permutations, p)
		}
	}
	TemplateRender(w, context.Filename, context)
}

