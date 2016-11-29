package main

import (
	"io"
	"net/http"

	"github.com/zyfdegh/fanach/dctl/api"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		io.WriteString(w, "Fanach dctl server")
	default:
		io.WriteString(w, "method not allowed")
	}
}

func handleDockerVersion(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.GetVersion(w, r)
	default:
		io.WriteString(w, "method not allowed")
	}
}
