package server

import (
	"github.com/gin-gonic/gin"
	"github.com/songjiayang/imagecloud/api/handler"
	"github.com/songjiayang/imagecloud/internal/config"
)

type Server struct {
	*gin.Engine
	*handler.Handler
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Engine:  gin.Default(),
		Handler: handler.NewHandler(cfg),
	}
}

func (s *Server) Listen() error {
	// register routes
	s.GET("/*key", func(c *gin.Context) {
		if c.Param("key") == "/" {
			s.Pong(c)
			return
		}

		s.Image.Get(c)
	})
	s.POST("/", s.Image.Post)

	return s.Run()
}
