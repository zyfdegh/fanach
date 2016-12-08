package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, resp interface{}) {
	body, err := json.Marshal(resp)
	if err != nil {
		log.Printf("marshal object error: %v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(body))
	return
}
