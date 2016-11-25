package service

import (
	"github.com/zyfdegh/fanach/deployer/server/entity"
	"github.com/zyfdegh/fanach/deployer/server/util"
)

// Deploy handles POST /deploy
func Deploy(reqDeploy *entity.ReqDeploy) (resp entity.RespPostDeploy, err error) {
	util.PrintPretty(reqDeploy, "reqDeploy")
	resp.Success = true
	return
}

func startSsContainer() {

}

func installDocker() {
	// cmd := "curl https://get.docker.com/ | sh"
	// executeCmd(config, host, port, cmd)
}
