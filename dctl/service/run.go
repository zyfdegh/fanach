package service

import (
	"log"
	"net/http"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/zyfdegh/fanach/dctl/entity"
	"github.com/zyfdegh/fanach/dctl/util"
)

const (
	defaultMem    = 32
	defaultCPU    = 0.01
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
	// Example docker run args
	// docker run -d --restart=always -p 8393:8388 zyfdedh/shadowsocks:latest -s 0.0.0.0 -p 8388 -k dev123 -m aes-256-cfb
	hostPort := req.HostPort
	if len(strings.TrimSpace(hostPort)) == 0 {
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

	mem := req.Mem
	if mem <= 0 {
		log.Printf("mem invalid, using default %d\n", defaultMem)
		mem = defaultMem
	}
	image := req.Image
	if strings.TrimSpace(image) == "" {
		log.Printf("image invalid, using default %s\n", defaultImage)
		image = defaultImage
	}
	cpu := req.CPU
	if cpu <= 0 {
		log.Printf("cpu invalid, using default %f\n", defaultCPU)
		cpu = defaultCPU
	}
	method := req.Method
	if !util.StringInSlice(method, supportedMethods) {
		log.Printf("unsupported method, using default %s\n", defaultMethod)
		method = defaultMethod
	}

	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		resp.Errmsg = err.Error()
		log.Printf("new docker client error: %v\n", err)
		return
	}

	cmd := []string{"-s", "0.0.0.0", "-p", "8388", "-k", password, "-m", method}

	hostConfig := &docker.HostConfig{
		PortBindings: map[docker.Port][]docker.PortBinding{
			"8388/tcp": []docker.PortBinding{
				docker.PortBinding{
					HostIP:   "0.0.0.0",
					HostPort: hostPort,
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

	container, err := client.CreateContainer(opts)
	if err != nil {
		resp.Errmsg = err.Error()
		log.Printf("docker create container error: %v\n", err)
		return
	}
	err = client.StartContainer(container.ID, hostConfig)
	if err != nil {
		// rm created container
		log.Printf("docker start container error: %v\n", err)
		log.Printf("remove container %s\n", container.ID)
		_, errRm := DockerRm(container.ID)
		if errRm != nil {
			resp.Errmsg = errRm.Error()
			log.Printf("docker remove container error: %v\n", errRm)
			return
		}

		resp.Errmsg = err.Error()
		log.Printf("docker start container error: %v\n", err)
		return
	}

	resp.Success = true
	resp.ID = container.ID
	return
}
