package service

import (
	"fmt"
	"log"

	"github.com/zyfdegh/fanach/deployer/server/entity"
	"golang.org/x/crypto/ssh"
)

// parse json config in request to SSHConfig
func parseConfig(reqDeploy *entity.ReqDeploy) (sshConfig *SSHConfig) {
	var clientConfig *ssh.ClientConfig
	if len(reqDeploy.Password) > 0 {
		clientConfig = passwdConfig(reqDeploy.User, reqDeploy.Password)
	}
	if len(reqDeploy.PrivateKey) > 0 {
		clientConfig = keyConfig(reqDeploy.User, reqDeploy.PrivateKey)
	}

	sshConfig = &SSHConfig{}
	sshConfig.ClientConfig = clientConfig
	sshConfig.Host = reqDeploy.Host
	sshConfig.Port = reqDeploy.Port
	return
}

func leaveFootprint(c *SSHConfig) (err error) {
	log.Println("logging date on server...")
	cmd := "date >> /tmp/fanach-deployer.log"
	output, err := executeRemote(c, cmd)
	if err != nil {
		log.Printf("execute cmd %s on server error: %v", cmd, err)
		return
	}
	fmt.Println(output)
	return
}
