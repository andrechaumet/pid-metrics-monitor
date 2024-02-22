package router

import (
	"github.com/gin-gonic/gin"
	"pid-metrics-monitor/controller"
	"pid-metrics-monitor/display"
)

const (
	BasePath = "/pmm/v1"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	pidRouter := r.Group("/pids")
	{
		pidRouter.POST(BasePath, controller.Create)
		pidRouter.PUT(BasePath, controller.Update)
		pidRouter.GET(BasePath, controller.FindAll)
	}
	go display.Display()
	return r
}
