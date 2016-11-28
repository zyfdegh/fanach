package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	startMockServer()
}

func startMockServer() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/api/ssaccount/", handleSsAccount)

	s := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	log.Printf("mock server start on localhost:%d", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start mock server error: %v", err)
	}
}
