package api

import (
	"net/http"
	"strings"

	"github.com/zyfdegh/fanach/dctl/entity"
	"github.com/zyfdegh/fanach/dctl/service"
)

// UnpauseContainer unpauses a docker container
func UnpauseContainer(w http.ResponseWriter, r *http.Request) {
	resp := &entity.Resp{}

	paramID := "id"
	id := r.FormValue(paramID)
	if strings.TrimSpace(id) == "" {
		resp.ErrNo = http.StatusBadRequest
		resp.Errmsg = "param id invalid"
		writeJSON(w, resp)
		return
	}

	err := service.DockerUnpause(id)
	if err != nil {
		resp.ErrNo = http.StatusInternalServerError
		resp.Errmsg = err.Error()
		writeJSON(w, resp)
		return
	}
	resp.Success = true
	writeJSON(w, resp)
	return
}
