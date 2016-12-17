package service

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

const (
	endpoint = "unix:///var/run/docker.sock"
)

var dockerClient *docker.Client

func newDockerClient(endpoint string) (*docker.Client, error) {
	c, err := docker.NewClient(endpoint)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func init() {
	c, err := newDockerClient(endpoint)
	if err != nil {
		log.Fatalf("new docker client error: %v\n", err)
	}
	if c != nil {
		dockerClient = c
	}
}
