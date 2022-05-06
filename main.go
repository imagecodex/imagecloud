package main

import (
	"log"

	"github.com/songjiayang/imgcloud/api/server"
	"github.com/songjiayang/imgcloud/internal/pkg/config"
)

func main() {
	cfg, err := config.Load("")
	if err != nil {
		log.Printf("load config with error; %v \n", err)
	}

	server.NewServer(cfg).Listen()
}
