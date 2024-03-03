package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pid-metrics-monitor/display"
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/service"
)

type PidDto struct {
	ID                int
	Name              string
	CurrentIterations int
	TotalIterations   int
	Logs              []string
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	pidRouter := r.Group("pmm/v1/pids")
	{
		pidRouter.POST("/", Create)
		pidRouter.PUT("/", Update)
		pidRouter.GET("/", FindAll)
	}
	go display.Display()
	return r
}

func Create(c *gin.Context) {
	var pid PidDto
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	c.JSON(http.StatusCreated, service.Create(toModel(pid)))
}

func Update(c *gin.Context) {
	var pid PidDto
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	service.Update(toModel(pid))
	c.Status(http.StatusAccepted)
}

func FindAll(c *gin.Context) {
	pids := service.FindAll()
	c.JSON(http.StatusOK, pids)
}

func bindAndValidate(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	return nil
}

func toModel(dto PidDto) model.Pid {
	var model model.Pid
	model.ID = dto.ID
	model.CurrentIterations = dto.CurrentIterations
	model.TotalIterations = dto.TotalIterations
	model.Name = dto.Name
	model.Logs = dto.Logs
	return model
}
