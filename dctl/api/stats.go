package api

import (
	"net/http"
	"strings"

	"github.com/zyfdegh/fanach/dctl/entity"
	"github.com/zyfdegh/fanach/dctl/service"
)

// GetDockerStats get usage infomation of a container
func GetDockerStats(w http.ResponseWriter, r *http.Request) {
	resp := &entity.RespGetStats{}

	paramID := "id"
	id := r.FormValue(paramID)
	if strings.TrimSpace(id) == "" {
		resp.ErrNo = http.StatusBadRequest
		resp.Errmsg = "param id invalid"
		writeJSON(w, resp)
		return
	}

	stats, err := service.DockerStats(id)
	if err != nil {
		resp.ErrNo = http.StatusInternalServerError
		resp.Errmsg = err.Error()
		writeJSON(w, resp)
		return
	}
	if stats != nil {
		resp.Stats = *stats
		resp.Success = true
	}
	writeJSON(w, resp)
	return
}
