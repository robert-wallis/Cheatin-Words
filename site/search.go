package site

import (
	"appengine"
	"fmt"
	"http"
	"log"
	"os"
	"strings"
	"word"
)

type Search struct {
	Filename     string
	Q            string
	Permutations []string
}

var enable *word.Enable
var enablePath = "enable.txt"

func searchHandler(w http.ResponseWriter, r *http.Request) {
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

	context := &Search{
		Filename:     "template/search.html",
		Q:            query,
		Permutations: make([]string, 0),
	}
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
	TemplateRender(w, context.Filename, context)
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
