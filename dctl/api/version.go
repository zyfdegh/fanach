package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/zyfdegh/fanach/dctl/service"
)

// GetDockerVersion queries version of docker on host machine
func GetDockerVersion(w http.ResponseWriter, r *http.Request) {
	resp, _ := service.Version()
	body, err := json.Marshal(resp)
	if err != nil {
		log.Printf("marshal object error: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(body))

	return
}
