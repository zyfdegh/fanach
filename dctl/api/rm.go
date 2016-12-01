package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/zyfdegh/fanach/dctl/entity"
	"github.com/zyfdegh/fanach/dctl/service"
)

// RmContainer deletes an existing docker container
func RmContainer(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body error: %v\n", err)
		return
	}
	var req entity.ReqPostRm
	err = json.Unmarshal(data, &req)
	if err != nil {
		log.Printf("unmarshal json error: %v\n", err)
		return
	}

	if len(strings.TrimSpace(req.ID)) == 0 {
		log.Printf("bad docker id %s\n", req.ID)
		return
	}

	resp, _ := service.DockerRm(req.ID)
	body, err := json.Marshal(resp)
	if err != nil {
		log.Printf("marshal object error: %v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(body))
	return
}
