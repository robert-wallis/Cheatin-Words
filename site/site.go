package site

import (
	"http"
)

func init() {
	http.HandleFunc("/", indexHandler)
}

