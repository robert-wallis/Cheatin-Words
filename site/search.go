package site

import (
	"http"
	"fmt"
	"scrabble"
)

type Search struct {
	Filename string
	Q string
	Permutations []string
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	if len(query) > 8 {
		query = query[0:8]
		http.Redirect(w, r, fmt.Sprintf("search?q=%s", query), 302)
	}
	context := &Search{
		Filename: "template/search.html",
		Q: query,
	}
	if 0 != len(query) {
		if len(query) > 8 {
			query = query[0:8]
		}
		channel := scrabble.StringPermutations(query)
		for p := range channel {
			context.Permutations = append(context.Permutations, p)
		}
	}
	TemplateRender(w, context.Filename, context)
}

// quickly calculate
func factorial(length int) int {
	result := length
	for i := length; i > 1; i-- {
		result *= i - 1
	}
	return result
}

