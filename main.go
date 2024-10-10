package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	http.HandleFunc("GET /", helloWorld)
	http.ListenAndServe(":"+port, nil)
}
