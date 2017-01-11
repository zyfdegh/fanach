package service

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

// DockerRm calls Docker API and rm container
func DockerRm(id string) (err error) {
	opts := docker.RemoveContainerOptions{}
	opts.ID = id
	opts.RemoveVolumes = true
	opts.Force = true
	err = dockerClient.RemoveContainer(opts)
	if err != nil {
		log.Printf("docker remove container error: %v\n", err)
		return
	}
	return
}
