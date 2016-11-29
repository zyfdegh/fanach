package service

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

// Version returns docker version
func Version() (ver *docker.Env, err error) {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Printf("new docker client error: %v\n", err)
		return
	}
	ver, err = client.Version()
	if err != nil {
		log.Printf("get docker version error: %v\n", err)
		return
	}
	log.Println(ver)
	return
}
