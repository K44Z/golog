package main

import (
	"fmt"
	"net/http"

	golog "github.com/K44Z/golog/internal"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(100)
	})
	loggedMux := golog.Log(mux)
	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", loggedMux)
}
