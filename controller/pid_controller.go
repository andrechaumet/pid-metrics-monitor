package controller

import (
	"net/http"
	//"pid-metrics-monitor/controller/dto"
	//"strconv"
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/service"
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

func Save(c *gin.Context) {
	var pid model.PidModel
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	service.Save(pid)
	c.JSON(statusCreated, gin.H{messageKey: "PID status created"})
}

func Update(c *gin.Context) {
	var pid model.PidModel
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	service.Update(pid)
	c.JSON(statusOK, gin.H{messageKey: "PID status updated"})
}

func FindAll(c *gin.Context) {
	pids := service.FindAll()
	c.JSON(statusOK, pids)
}

func bindAndValidate(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(badRequest, gin.H{errorKey: err.Error()})
		return err
	}
	return nil
}