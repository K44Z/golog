package main

import (
	"fmt"
	"net/http"

	gologger "github.com/K44Z/gologger/internal"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "wassup")
	})
	loggedMux := gologger.Log(mux)
	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", loggedMux)
}
