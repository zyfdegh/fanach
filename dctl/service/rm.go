package service

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/zyfdegh/fanach/dctl/entity"
)

// DockerRm calls Docker API and rm container
func DockerRm(id string) (resp entity.Resp, err error) {
	opts := docker.RemoveContainerOptions{}
	opts.ID = id
	opts.RemoveVolumes = true
	opts.Force = true
	err = dockerClient.RemoveContainer(opts)
	if err != nil {
		resp.Errmsg = err.Error()
		log.Printf("docker remove container error: %v\n", err)
		return
	}
	resp.Success = true
	return
}
