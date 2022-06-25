package handler

import (
	"github.com/go-kit/log"

	"github.com/songjiayang/imagecloud/internal/config"
)

type Handler struct {
	*Image
	*Ponger
	logger log.Logger
}

func NewHandler(cfg *config.Config, logger log.Logger) *Handler {
	return &Handler{
		Image: &Image{
			enableSites: cfg.EnableSites,
			logger:      logger,
		},
		Ponger: new(Ponger),
	}
}
