package main

import (
	"fmt"
	"net/http"
)

func createServer() {
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.Handle("/assets", http.FileServer(http.Dir("./assets/logo.png")))
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
