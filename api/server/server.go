package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/songjiayang/imagecloud/api/handler"
	"github.com/songjiayang/imagecloud/internal/config"
)

type Server struct {
	*gin.Engine
	*handler.Handler
	logger log.Logger
}

func NewServer(cfg *config.Config, logger log.Logger) *Server {
	return &Server{
		Engine:  gin.Default(),
		Handler: handler.NewHandler(cfg, logger),
	}
}

func (s *Server) Listen() error {
	s.routes()

	return s.Run()
}

func (s *Server) routes() {
	s.GET("/*key", func(c *gin.Context) {
		if c.Param("key") == "/" {
			s.Pong(c)
			return
		}

		s.Image.Get(c)
	})

	s.POST("/", s.Image.Post)
}
