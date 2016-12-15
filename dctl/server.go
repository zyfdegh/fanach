package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func startServer() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/dver", handleDockerVersion)
	http.HandleFunc("/ssc", handleSsContainer)
	http.HandleFunc("/stats", handleDockerStats)

	s := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	log.Printf("server start on localhost:%d\n", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start server error: %v\n", err)
	}
}
