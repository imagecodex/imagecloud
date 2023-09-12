package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/songjiayang/imagecloud/internal/metrics"
)

func Duration(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()

	path := ctx.Request.URL.Path
	if path != "/" && path != "/metrics" {
		metrics.HttpRequestDuration.WithLabelValues(
			ctx.Request.Method,
			strconv.Itoa(ctx.Writer.Status()),
		).Observe(time.Since(start).Seconds())
	}
}
