package router

import (
   "github.com/gin-gonic/gin"
   "pid-metrics-monitor/controller"
)

func SetupRouter() *gin.Engine {
   r := gin.Default()
   pidRouter := r.Group("/pids")
   {
      pidRouter.POST("", controller.CreatePID)
      pidRouter.PUT("/:id", controller.UpdatePID)
      pidRouter.POST("/:id/logs", controller.CreatePIDLog)
   }
   return r
}