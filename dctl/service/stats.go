package service

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/zyfdegh/fanach/dctl/entity"
)

// DockerStats calls Docker API to get container stats
// TODO Optimize me, it takes over 1000ms to stat a container
func DockerStats(id string) (stats *entity.Stats, err error) {
	stats = &entity.Stats{}

	errCh := make(chan error, 1)
	statsCh := make(chan *docker.Stats)
	doneCh := make(chan bool)
	defer close(doneCh)
	// docker client will close this chan automatically
	// defer close(statsCh)

	go func() {
		errCh <- dockerClient.Stats(docker.StatsOptions{ID: id, Stats: statsCh, Stream: false, Done: doneCh})
		close(errCh)
	}()

	var resultStats []docker.Stats
	for {
		s, ok := <-statsCh
		if !ok {
			break
		}
		resultStats = append(resultStats, *s)
	}

	err = <-errCh
	if err != nil {
		log.Printf("stats container error: %v\n", err)
		return
	}

	if len(resultStats) != 1 {
		log.Printf("len result stats[%d] is not 1\n", len(resultStats))
		return
	}

	for _, ifStats := range resultStats[0].Networks {
		stats.RxBytes += ifStats.RxBytes
		stats.TxBytes += ifStats.TxBytes
	}

	return
}
