package router

import (
   "github.com/gin-gonic/gin"
   "pid-metrics-monitor/controller"
)

const (
	BasePath  = ""
)

func SetupRouter() *gin.Engine {
   r := gin.Default()
   pidRouter := r.Group("/pids")
   {
      pidRouter.POST(BasePath, controller.Save)
      pidRouter.PUT(BasePath, controller.Update)
      pidRouter.GET(BasePath, controller.FindAll)
   }
   return r
}