package handler

import (
	"net/http"
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/service"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	statusCreated       = http.StatusCreated
	statusOK            = http.StatusOK
	badRequest          = http.StatusBadRequest
	messageKey          = "message"
	errorKey            = "error"
	defaultErrorMessage = "Something went wrong"
)

type PidDto struct {
	ID                int
	StartTime         time.Time
	CurrentIterations int
	TotalIterations   int
	Logs 			  []string
}

func Save(c *gin.Context) {
	var pid PidDto
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	service.Save(toModel(pid))
	c.JSON(statusCreated, gin.H{messageKey: "PID created"})
}

func Update(c *gin.Context) {
	var pid PidDto
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	service.Update(toModel(pid))
	c.JSON(statusOK, gin.H{messageKey: "PID updated"})
}

func bindAndValidate(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(badRequest, gin.H{errorKey: err.Error()})
		return err
	}
	return nil
}

func FindAll(c *gin.Context) {
	pids := service.FindAll()
	c.JSON(statusOK, pids)
}

func toModel(dto PidDto) model.PidModel {
	var model model.PidModel
	model.CurrentIterations = dto.CurrentIterations
	model.TotalIterations = dto.TotalIterations
	model.Logs = dto.Logs
	return model
}