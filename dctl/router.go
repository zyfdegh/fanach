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
	defer r.Body.Close()
}

func handleDockerVersion(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.GetVersion(w, r)
	default:
		io.WriteString(w, "method not allowed")
	}
	defer r.Body.Close()
}

func handleSsContainer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		api.AddSsContainer(w, r)
	case http.MethodDelete:
		api.RmContainer(w, r)
	default:
		io.WriteString(w, "method not allowed")
	}
	defer r.Body.Close()
}

func handlePauseContainer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		api.PauseContainer(w, r)
	case http.MethodDelete:
		api.UnpauseContainer(w, r)
	default:
		io.WriteString(w, "method not allowed")
	}
	defer r.Body.Close()
}

func handleDockerStats(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.GetDockerStats(w, r)
	default:
		io.WriteString(w, "method not allowed")
	}
	defer r.Body.Close()
}
