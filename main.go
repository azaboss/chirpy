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
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	hh := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
	mux.HandleFunc("/healthz", hh)
	s.ListenAndServe()
}
