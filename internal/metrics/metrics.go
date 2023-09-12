package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	namespace           = "imagecloud"
	HttpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "http_duration_seconds",
		Help:      "Total http request duration in seconds",
	}, []string{"method", "code"})

	ImageRemoteLoadDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "image_remote_load_duration_seconds",
		Help:      "Total remote image load duration in seconds",
	})
)

func NewHandler() http.Handler {
	reg := prometheus.NewRegistry()
	registerMetrics(reg)
	return promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		Registry:          reg,
		EnableOpenMetrics: true,
	})
}

func registerMetrics(reg prometheus.Registerer) {
	reg.MustRegister(
		HttpRequestDuration,
		ImageRemoteLoadDuration,
	)
}
