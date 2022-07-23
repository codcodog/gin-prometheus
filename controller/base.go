package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Demo(ctx *gin.Context) {
	data := map[string]string{
		"hello": "world",
	}

	ctx.JSON(http.StatusOK, data)
	return
}

func (c *Controller) Duration(ctx *gin.Context) {
	time.Sleep(3 * time.Second)

	data := map[string]string{
		"duration": "sleep 3 seconds",
	}
	ctx.JSON(http.StatusOK, data)
	return
}

func (c *Controller) Code(ctx *gin.Context) {
	data := map[string]string{
		"code": "http code is 400",
	}

	ctx.JSON(http.StatusBadRequest, data)
	return
}
