package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Setting up static file server
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
