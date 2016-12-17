package service

import "log"

// DockerPause pauses a container
func DockerPause(id string) (err error) {
	err = dockerClient.PauseContainer(id)
	if err != nil {
		log.Printf("pause container error: %v\n", err)
		return
	}
	return
}
