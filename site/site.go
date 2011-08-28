package site

import (
	"http"
	"scrabble"
)

func init() {
	http.HandleFunc("/", indexHandler)
}

type Index struct {
	Permutations []string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	context := &Index{}
	if 0 != len(query) {
		context.Permutations = make([]string, factorial(len(query)))
		channel := scrabble.StringPermutations(query)
		i := 0
		for p := range channel {
			context.Permutations[i] = p
			i++
		}
	}
	renderTemplate(w, "template/index.html", context)
}

// quickly calculate
func factorial(length int) (int) {
	result := length
	for i := length; i > 1; i-- {
		result *= i-1
	}
	return result;
}

func renderTemplate(w http.ResponseWriter, filename string, context interface{}) {
	page := Page{}
	template, err := page.Load(filename)
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
	if err := template.Execute(w, context); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}

