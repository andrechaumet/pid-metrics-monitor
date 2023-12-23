package router

import (
   "github.com/gin-gonic/gin"
   "pid-metrics-monitor/controller"
)

func SetupRouter() *gin.Engine {
   r := gin.Default()
   pidRouter := r.Group("/pids")
   {
      pidRouter.POST("", controller.Save)
      pidRouter.PUT("", controller.Update)
      pidRouter.GET("", controller.FindAll)
      //pidRouter.POST("/:id/logs", controller.AddPidLog)
   }
   return r
}