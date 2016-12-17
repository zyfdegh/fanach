package service

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

// DockerUnpause unpauses a container
func DockerUnpause(id string) (err error) {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Printf("new docker client error: %v\n", err)
		return
	}
	err = client.UnpauseContainer(id)
	if err != nil {
		log.Printf("unpause container error: %v\n", err)
		return
	}
	return
}
