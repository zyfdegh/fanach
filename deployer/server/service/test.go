package service

import (
	"log"

	"github.com/zyfdegh/fanach/deployer/server/entity"
)

// TestConnection check if server can be connected
func TestConnection(reqDeploy *entity.ReqDeploy) (resp entity.Resp, err error) {
	c := parseConfig(reqDeploy)
	err = leaveFootprint(c)
	if err != nil {
		log.Printf("leave footprint error: %v", err)
		resp.Success = false
		resp.Errmsg = err.Error()
		return
	}

	resp.Success = true
	return
}
