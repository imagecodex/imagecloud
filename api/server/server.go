package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"

	"github.com/imagecodex/imagecloud/api/handler"
	"github.com/imagecodex/imagecloud/internal/config"
	"github.com/imagecodex/imagecloud/internal/metrics"
	"github.com/imagecodex/imagecloud/internal/middleware"
)

type Server struct {
	*gin.Engine
	*handler.Handler
}

func NewServer(cfg *config.Config, logger log.Logger) *Server {
	engine := gin.Default()
	engine.Use(middleware.Duration)
	return &Server{
		Engine:  engine,
		Handler: handler.NewHandler(cfg, logger),
	}
}

func (s *Server) Listen() error {
	s.routes()

	return s.Run()
}

func (s *Server) routes() {
	metricsHandler := metrics.NewHandler()

	s.GET("/*key", func(c *gin.Context) {
		switch c.Param("key") {
		case "/":
			s.Pong(c)
		case "/metrics":
			metricsHandler.ServeHTTP(c.Writer, c.Request)
		default:
			s.Image.Get(c)
		}
	})

	s.POST("/*key", s.Image.Post)
}
