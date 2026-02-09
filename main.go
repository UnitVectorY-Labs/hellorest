package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime/debug"
)

// Version is the application version, injected at build time via ldflags
var Version = "dev"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-App-Version", Version)
	json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
}

func main() {
	// Set the build version from the build info if not set by the build system
	if Version == "dev" || Version == "" {
		if bi, ok := debug.ReadBuildInfo(); ok {
			if bi.Main.Version != "" && bi.Main.Version != "(devel)" {
				Version = bi.Main.Version
			}
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	http.HandleFunc("GET /", helloWorld)
	http.ListenAndServe(":"+port, nil)
}
