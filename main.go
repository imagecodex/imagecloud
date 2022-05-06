package main

import (
	"log"

	"github.com/songjiayang/imagecloud/api/server"
	"github.com/songjiayang/imagecloud/internal/pkg/config"
)

func main() {
	cfg, err := config.Load("")
	if err != nil {
		log.Printf("load config with error; %v \n", err)
	}

	server.NewServer(cfg).Listen()
}
