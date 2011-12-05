package site

import (
	"appengine"
	"fmt"
	"http"
	"log"
	"os"
	"strings"
	"template"
	"word"
)

type Search struct {
	T			*template.Template
	Q            string
	Permutations []string
}

var enable *word.Enable
var enablePath = "enable.txt"

func searchHandler(w http.ResponseWriter, r *http.Request) {
	context := &Search{}
	// prepare the template
	t, err := template.ParseFile("template/search.html")
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
	context.T = t
	ae := appengine.NewContext(r)
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
	context.Q = query

	hashTable := make(map[string]byte, 0)

	if 0 != queryLength {
		if queryLength > 8 {
			query = query[0:8]
		}
		// make sure enable is loaded
		var e *word.Enable
		var err os.Error
		if e, err = loadEnable(); err != nil {
			ae.Errorf("%v", err)
		}
		channel := word.StringPermutations(query)
		for p := range channel {
			if valid := e.WordIsValid(p); !valid {
				continue
			}
			if _, inHash := hashTable[p]; !inHash {
				context.Permutations = append(context.Permutations, p)
				hashTable[p] = 1
			}
		}
	}
	// display template
	if err := context.T.Execute(w, context); err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}

// ensure this instance has the Enable dictionary loaded
func loadEnable() (*word.Enable, os.Error) {
	if enable != nil {
		return enable, nil
	}
	log.Println("Loading dictionary, this query should be .3 seconds slower")
	enable = new(word.Enable)
	if err := enable.Load(enablePath); err != nil {
		return nil, err
	}
	return enable, nil
}
