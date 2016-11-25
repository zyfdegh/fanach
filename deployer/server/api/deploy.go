package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zyfdegh/fanach/deployer/server/entity"
	"github.com/zyfdegh/fanach/deployer/server/service"
)

// GetRoot serves index.html
func GetRoot(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/html/index.html")
}

// PostDeploy starts deployment
func PostDeploy(w http.ResponseWriter, req *http.Request) {
	var reqDeploy = &entity.ReqDeploy{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("read request error: %v", err)
		return
	}
	err = json.Unmarshal(data, reqDeploy)
	if err != nil {
		log.Printf("unmarshal to object error: %v", err)
		return
	}
	resp, err := service.Deploy(reqDeploy)
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
