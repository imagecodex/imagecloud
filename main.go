package main

import (
	"flag"
	"os"

	"github.com/go-kit/log"

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
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	cfg, err := config.Load(cfgPath)
	if err != nil {
		logger.Log("msg", "load config failed", "error", err)
		return
	}

	if err := server.NewServer(cfg, logger).Listen(); err != nil {
		logger.Log("msg", "server start failed", "error", err)
	}
}
