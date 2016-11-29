package server

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

// Start launches http server
func Start() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/test", handleTest)
	http.HandleFunc("/deploy", handleDeploy)

	s := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	log.Printf("server start on localhost:%d", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start server error: %v", err)
	}
}
