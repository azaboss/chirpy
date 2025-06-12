package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Starting server on :8080")
	s.ListenAndServe()
}
