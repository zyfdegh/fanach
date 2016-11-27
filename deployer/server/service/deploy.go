package service

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"

	"github.com/zyfdegh/fanach/deployer/server/entity"
)

// Deploy handles POST /deploy
func Deploy(reqDeploy *entity.ReqDeploy) (resp entity.RespPostDeploy, err error) {
	resp.Success = true
	var clientConfig *ssh.ClientConfig
	if len(reqDeploy.Password) > 0 {
		clientConfig = passwdConfig(reqDeploy.User, reqDeploy.Password)
	}
	if len(reqDeploy.PrivateKey) > 0 {
		clientConfig = keyConfig(reqDeploy.User, reqDeploy.PrivateKey)
	}

	c := &SSHConfig{}
	c.ClientConfig = clientConfig
	c.Host = reqDeploy.Host
	c.Port = reqDeploy.Port

	leaveFootprints(c)

	return
}

func leaveFootprints(c *SSHConfig) {
	log.Println("logging date on server...")
	cmd := "date >> /tmp/fanach-deployer.log"
	output, err := executeRemote(c, cmd)
	if err != nil {
		log.Printf("execute cmd %s on server error: %v", cmd, err)
		return
	}
	fmt.Println(output)
}

func startSsContainer() {

}

func installDocker() {
	// cmd := "curl https://get.docker.com/ | sh"
	// executeCmd(config, host, port, cmd)
}

// startup docker on system boot and launch docker now
func enableDocker() {

}
