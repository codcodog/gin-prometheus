package main

import (
	"codcodog/ginprometheus/controller"
	"codcodog/ginprometheus/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.New()
	c := controller.New()

	go func() {
		prometheus.Register(middleware.ReqCounter)
		prometheus.Register(middleware.ReqDuration)

		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8081", nil)
	}()

	r.Use(middleware.Prometheus())
	r.GET("/demo", c.Demo)
	r.GET("/code", c.Code)
	r.POST("/duration", c.Duration)

	r.Run(":8080")
}
