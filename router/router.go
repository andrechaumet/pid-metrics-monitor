package router

import (
   "github.com/gin-gonic/gin"
   "pid-metrics-monitor/handler"
   "pid-metrics-monitor/batch"
)

const (
	BasePath  = ""
)

func SetupRouter() *gin.Engine {
   r := gin.Default()
   pidRouter := r.Group("/pids")
   {
      pidRouter.POST(BasePath, handler.Save)
      pidRouter.PUT(BasePath, handler.Update)
      pidRouter.GET(BasePath, handler.FindAll)
   }
   go batch.LoadAndDisplay()
   return r
}