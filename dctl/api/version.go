package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/zyfdegh/fanach/dctl/service"
)

// GetVersion queries version of docker on host machine
func GetVersion(w http.ResponseWriter, r *http.Request) {
	resp, err := service.Version()
	if err != nil {
		log.Printf("serve deploy error: %v", err)
		return
	}
	body, err := json.Marshal(resp)
	if err != nil {
		log.Printf("marshal object error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(body))

	return
}
