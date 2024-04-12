package main

import (
	"fmt"
	"net/http"

	"github.com/hs3city/tidyshelf/internal/app"
)

func main() {
	router := http.NewServeMux()

	_ = &app.App{
		// Model: gemini.Gemini{},
		Model: struct{}{},
	}
	router.HandleFunc("POST /bookshelf", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("", "")

		fmt.Fprintf(w, "")
	})

	http.ListenAndServe(":8080", router)
}
