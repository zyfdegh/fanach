package service

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

// DockerPause pauses a container
func DockerPause(id string) (err error) {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Printf("new docker client error: %v\n", err)
		return
	}
	err = client.PauseContainer(id)
	if err != nil {
		log.Printf("pause container error: %v\n", err)
		return
	}
	return
}
