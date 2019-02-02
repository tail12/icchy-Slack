package main

import (
	"log"
	"net/http"
)

func interaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("error: Invalid method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("test")
}
