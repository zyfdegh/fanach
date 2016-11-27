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
	installDocker(c)
	enableDocker(c)
	startSsContainer(c)

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

func installDocker(c *SSHConfig) {
	log.Println("installing docker on server...")
	cmd := "curl https://get.docker.com/ | sh"
	output, err := executeRemote(c, cmd)
	if err != nil {
		log.Printf("execute cmd %s on server error: %v", cmd, err)
		return
	}
	fmt.Println(output)
}

// startup docker on system boot and launch docker now
func enableDocker(c *SSHConfig) {
	log.Println("enabling docker on server...")
	cmd := "systemctl enable docker && systemctl start docker"
	output, err := executeRemote(c, cmd)
	if err != nil {
		log.Printf("execute cmd %s on server error: %v", cmd, err)
		return
	}
	fmt.Println(output)
}

func startSsContainer(c *SSHConfig) {
	log.Println("starting a ss container on server...")
	cmd := "docker run -d -p 8387:8388 zyfdedh/shadowsocks:latest -s 0.0.0.0 -p 8388 -k deployer123 -m aes-256-cfb"
	output, err := executeRemote(c, cmd)
	if err != nil {
		log.Printf("execute cmd %s on server error: %v", cmd, err)
		return
	}
	fmt.Println(output)
}
