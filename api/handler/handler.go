package handler

import "github.com/songjiayang/imgcloud/internal/pkg/config"

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
