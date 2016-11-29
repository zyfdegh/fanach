package service

import (
	"fmt"
	"log"

	"github.com/zyfdegh/fanach/deployer/server/entity"
)

// Deploy handles POST /deploy
func Deploy(reqDeploy *entity.ReqDeploy) (resp entity.RespPostDeploy, err error) {
	c := parseConfig(reqDeploy)

	leaveFootprint(c)
	installDocker(c)
	enableDocker(c)
	startSsContainer(c)

	resp.Success = true
	return
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
