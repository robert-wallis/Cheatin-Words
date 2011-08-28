package site

import (
	"http"
	"scrabble"
)

type Search struct {
	Filename string
	Q string
	Permutations []string
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	context := &Search{
		Filename: "template/search.html",
		Q: query,
	}
	if 0 != len(query) {
		context.Permutations = make([]string, factorial(len(query)))
		channel := scrabble.StringPermutations(query)
		i := 0
		for p := range channel {
			context.Permutations[i] = p
			i++
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

