package handler

import (
	"net/http"
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/service"
	"time"

	"github.com/gin-gonic/gin"
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
	Logs              []string
}

func Create(c *gin.Context) {
	var pid PidDto
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	service.Create(toModel(pid))
	c.Status(http.StatusCreated)
}

func Save(c *gin.Context) {
	var pid PidDto
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	service.Save(toModel(pid))
	c.Status(http.StatusCreated)
}

func Update(c *gin.Context) {
	var pid PidDto
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	service.Update(toModel(pid))
	c.Status(http.StatusAccepted)
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
