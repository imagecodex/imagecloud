package main

import (
	"flag"
	"log"
	"os"

	kitlog "github.com/go-kit/log"
	"github.com/songjiayang/imagecloud/api/server"
	"github.com/songjiayang/imagecloud/internal/config"
)

var (
	cfgPath string
)

func init() {
	flag.StringVar(&cfgPath, "config.file", "", "imagecloud config file path")
	flag.Parse()
}

func main() {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	log.SetOutput(kitlog.NewStdlibAdapter(logger))

	cfg, err := config.Load(cfgPath)
	if err != nil {
		log.Printf("load config with error; %v", err)
	}

	server.NewServer(cfg).Listen()
}
