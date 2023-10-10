package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handle requests for the index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Handle requests for the second page
	http.HandleFunc("/second.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "second.html")
	})

	// Start the web server on port 8080
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
