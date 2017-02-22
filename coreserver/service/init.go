package service

import "log"

func init() {
	if err := initUserDB(); err != nil {
		log.Printf("error init user db: %v\n", err)
	}
}
