package service

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/zyfdegh/fanach/dctl/entity"
	"github.com/zyfdegh/fanach/dctl/util"
)

const (
	defaultImage  = "zyfdedh/shadowsocks:latest"
	defaultMethod = "aes-256-cfb"
)

var (
	supportedMethods = []string{
		"aes-128-cfb",
		"aes-192-cfb",
		"aes-256-cfb",
		"des-cfb",
		"bf-cfb",
		"cast5-cfb",
		"rc4-md5",
		"chacha20",
		"salsa20",
	}
)

// DockerRun starts a new ss container
func DockerRun(req entity.ReqPostRun) (resp entity.RespPostRun, err error) {
	hostPort := req.HostPort
	if hostPort <= 0 {
		resp.Errmsg = "hostPort not set"
		resp.ErrNo = http.StatusBadRequest
		return
	}
	password := req.Password
	if len(strings.TrimSpace(password)) == 0 {
		resp.Errmsg = "password not set"
		resp.ErrNo = http.StatusBadRequest
		return
	}

	image := req.Image
	if strings.TrimSpace(image) == "" {
		image = defaultImage
	}
	method := req.Method
	if !util.StringInSlice(method, supportedMethods) {
		method = defaultMethod
	}
	mem := req.Mem
	cpu := req.CPU

	cmd := []string{"-s", "0.0.0.0", "-p", "8388", "-k", password, "-m", method}

	hostConfig := &docker.HostConfig{
		PortBindings: map[docker.Port][]docker.PortBinding{
			"8388/tcp": []docker.PortBinding{
				docker.PortBinding{
					HostIP:   "0.0.0.0",
					HostPort: fmt.Sprintf("%d", hostPort),
				},
			},
		},
	}
	opts := docker.CreateContainerOptions{
		Config: &docker.Config{
			Image:     image,
			Memory:    int64(mem * 1024 * 1024),
			CPUShares: int64(cpu * 1024),
			Cmd:       cmd,
			ExposedPorts: map[docker.Port]struct{}{
				"8388/tcp": {},
			},
		},
		HostConfig: hostConfig,
	}

	container, err := dockerClient.CreateContainer(opts)
	if err != nil {
		resp.Errmsg = err.Error()
		log.Printf("docker create container error: %v\n", err)
		return
	}
	err = dockerClient.StartContainer(container.ID, hostConfig)
	if err != nil {
		log.Printf("docker start container error: %v\n", err)
		resp.Errmsg = err.Error()

		// rm created container
		log.Printf("remove container %s\n", container.ID)
		errRm := DockerRm(container.ID)
		if errRm != nil {
			resp.Errmsg = errRm.Error()
			log.Printf("docker remove container error: %v\n", errRm)
			return
		}

		return
	}

	resp.Success = true
	resp.ID = container.ID
	return
}
