package service

import "log"

// DockerUnpause unpauses a container
func DockerUnpause(id string) (err error) {
	err = dockerClient.UnpauseContainer(id)
	if err != nil {
		log.Printf("unpause container error: %v\n", err)
		return
	}
	return
}
