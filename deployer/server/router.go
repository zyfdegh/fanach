package server

import (
	"io"
	"net/http"

	"github.com/zyfdegh/fanach/deployer/server/api"
)

func handleRoot(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		api.GetRoot(w, req)
	default:
		io.WriteString(w, "method not allowed")
	}
}

func handleDeploy(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		api.PostDeploy(w, req)
	default:
		io.WriteString(w, "method not allowed")
	}
}
