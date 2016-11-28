package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		io.WriteString(w, "Fanach mock server")
	default:
		io.WriteString(w, "method not allowed")
	}
}

func handleSsAccount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		resp, _ := getSsAccount()
		body, err := json.Marshal(resp)
		if err != nil {
			log.Printf("marshal object error: %v", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(body))
	case http.MethodPost:
	default:
		io.WriteString(w, "method not allowed")
	}
}
