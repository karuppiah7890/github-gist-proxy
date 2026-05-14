package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running GitHub Gist Proxy")

	mux := http.NewServeMux()

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("error occurred while trying to run server: %v", err)
	}
}
