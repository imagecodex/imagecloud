package handler

import "github.com/songjiayang/imagecloud/internal/pkg/config"

type Handler struct {
	Image *Image
	*Ponger
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		Image: &Image{
			enableSites: cfg.EnableSites,
		},
		Ponger: &Ponger{},
	}
}
