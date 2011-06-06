package hello

import (
	"fmt"
	"http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.FormValue("yourmom"))
}
