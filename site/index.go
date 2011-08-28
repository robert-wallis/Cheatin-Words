package site

import (
	"http"
	"scrabble"
)

type Index struct {
	filename string
	Permutations []string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	context := &Index{ filename:"template/index.html" }
	if 0 != len(query) {
		context.Permutations = make([]string, factorial(len(query)))
		channel := scrabble.StringPermutations(query)
		i := 0
		for p := range channel {
			context.Permutations[i] = p
			i++
		}
	}
	TemplateRender(w, context.filename, context)
}

// quickly calculate
func factorial(length int) int {
	result := length
	for i := length; i > 1; i-- {
		result *= i - 1
	}
	return result
}

