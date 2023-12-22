package controller

import (
	"net/http"
	"pid-metrics-monitor/controller/dto"
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/service"
	"strconv"
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

func CreatePid(c *gin.Context) {
	var pidModel model.PidModel
	if err := bindAndValidate(c, &pidModel); err != nil {
		return
	}
	service.Save(pidModel)
	c.JSON(statusCreated, gin.H{messageKey: "PID status created"})
}

func UpdatePid(c *gin.Context) {
	var pidModel model.PidModel
	if err := bindAndValidate(c, &pidModel); err != nil {
		return
	}
	service.Update(pidModel)
	c.JSON(statusOK, gin.H{messageKey: "PID status updated"})
}

func AddPidLog(c *gin.Context) {
	pidID := c.Param("id")
	id, err := strconv.Atoi(pidID)
	if err != nil {
		c.JSON(badRequest, gin.H{errorKey: "Invalid ID"})
		return
	}

	var log dto.LogRequest
	if err := bindAndValidate(c, &log); err != nil {
		return
	}

	service.AddPidLog(id, log.LogMessage)
	c.JSON(statusCreated, gin.H{messageKey: "Log added"})
}

func bindAndValidate(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(badRequest, gin.H{errorKey: err.Error()})
		return err
	}
	return nil
}