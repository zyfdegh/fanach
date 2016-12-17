package service

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

// Version returns docker version
func Version() (ver *docker.Env, err error) {
	ver, err = dockerClient.Version()
	if err != nil {
		log.Printf("get docker version error: %v\n", err)
		return
	}
	return
}
