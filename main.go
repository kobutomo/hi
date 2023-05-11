package main

import (
	"fmt"
	"log"
	"net/http"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/v1/hi requested")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "{\"message\": \"hi there\"}")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("/v1/healthcheck requested")
	w.WriteHeader(http.StatusOK)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/v1/hi", hiHandler)
	http.HandleFunc("/v1/healthcheck", healthCheck)
	port := "8080"

	log.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %s\n", err)
	}
}
