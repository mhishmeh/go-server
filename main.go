package main

import (
	"fmt"
	"net/http"
)

func createServer() {
	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	addr := server.Addr
	// Start the server
	fmt.Println("Starting server on", addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {

	createServer()
}
