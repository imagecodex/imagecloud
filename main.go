package main

import (
	"flag"
	"log"

	"github.com/songjiayang/imagecloud/api/server"
	"github.com/songjiayang/imagecloud/internal/pkg/config"
)

var (
	cfgPath string
)

func init() {
	flag.StringVar(&cfgPath, "config.file", "", "imagecloud config file path")
	flag.Parse()
}

func main() {
	cfg, err := config.Load(cfgPath)
	if err != nil {
		log.Printf("load config with error; %v \n", err)
	}

	server.NewServer(cfg).Listen()
}
