package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ponger struct {
}

func (*Ponger) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
