package main

import (
	"net/http"

	"github.com/dslcosta1/anemone/view/templates"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Index().Render(r.Context(), w)
	})

	println("Frontend running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}