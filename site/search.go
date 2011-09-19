package site

import (
	"fmt"
	"http"
	"log"
	"strings"
	"word"
)

type Search struct {
	Filename     string
	Q            string
	Permutations []string
}

var enable *word.Enable
var enablePath = "static/enable.txt"

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	queryLength := len([]int(query)) // unicode length
	// max 8 letters
	if queryLength > 8 {
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
		Q:        query,
	}

	if 0 != queryLength {
		if queryLength > 8 {
			query = query[0:8]
		}
		// make sure enable is loaded
		e := loadEnable()
		channel := word.StringPermutations(query)
		for p := range channel {
			if !e.WordIsValid(p) {
				continue
			}
			context.Permutations = append(context.Permutations, p)
		}
	}
	TemplateRender(w, context.Filename, context)
}

// ensure this instance has the Enable dictionary loaded
func loadEnable() *word.Enable {
	if enable != nil {
		return enable
	}
	log.Println("Loading dictionary, this query should be .3 seconds slower")
	enable = new(word.Enable)
	enable.Load(enablePath)
	return enable
}
