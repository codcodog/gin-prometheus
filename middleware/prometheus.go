package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var ReqCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	},
	[]string{"code", "method", "url"},
)

var ReqDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "request_duration_seconds",
		Help: "The HTTP request latencies in seconds.",
	},
	[]string{"code", "method", "url"},
)

func Prometheus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		duration := time.Since(start).Seconds()
		code := strconv.Itoa(ctx.Writer.Status())
		method := ctx.Request.Method
		url := ctx.Request.URL.Path

		ReqCounter.WithLabelValues(code, method, url).Inc()
		ReqDuration.WithLabelValues(code, method, url).Observe(duration)
	}
}
