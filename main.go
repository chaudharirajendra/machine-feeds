package main

import (
	"fmt"
	"machine-feeds/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/machine-feeds", handlers.MachineFeedHandler)

	port := 8080
	fmt.Printf("Server running at http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
