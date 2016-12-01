package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zyfdegh/fanach/dctl/entity"
	"github.com/zyfdegh/fanach/dctl/service"
)

// AddSsContainer starts a new ss container
func AddSsContainer(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body error: %v\n", err)
		return
	}

	var req entity.ReqPostRun
	err = json.Unmarshal(data, &req)
	if err != nil {
		log.Printf("unmarshal json error: %v\n", err)
		return
	}

	resp, _ := service.DockerRun(req)
	body, err := json.Marshal(resp)
	if err != nil {
		log.Printf("marshal object error: %v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(body))
	return
}
