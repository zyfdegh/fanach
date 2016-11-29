package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func startServer() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/docker/version", handleDockerVersion)

	s := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	log.Printf("server start on localhost:%d", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start server error: %v", err)
	}
}
