package api

import (
	"io"
	"net/http"
)

// GetRoot handles GET /
func GetRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Fanach dctl server")
}
