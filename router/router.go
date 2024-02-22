package router

import (
	"github.com/gin-gonic/gin"
	"pid-metrics-monitor/handler"
	"pid-metrics-monitor/monitor"
)

const (
	BasePath = "/pmm/v1"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	pidRouter := r.Group("/pids")
	{
		pidRouter.POST(BasePath, handler.Create)
		pidRouter.PUT(BasePath, handler.Update)
		pidRouter.GET(BasePath, handler.FindAll)
	}
	go monitor.Display()
	return r
}
