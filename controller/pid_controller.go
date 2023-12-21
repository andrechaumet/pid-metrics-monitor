package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pid-metrics-monitor/service"
	"pid-metrics-monitor/dto"
	"pid-metrics-monitor/model"
	"strconv"
 )
 
 func CreatePid(c *gin.Context) {
	var pidModel model.PidModel
	if err := validatePidBody(c, &pidModel); err != nil {
		return
	}
	service.Save(pidModel)
	c.JSON(http.StatusCreated, gin.H{"message": "PID status created"})
}
 
 func UpdatePid(c *gin.Context) {
	var pidModel model.PidModel
	if err := validatePidBody(c, &pidModel); err != nil {
		return
	}
	service.Update(pidModel)
	c.JSON(http.StatusOK, gin.H{"message": "PID status updated"})
 }
 
 func AddPidLog(c *gin.Context) {
	pidID := c.Param("id")
	id, _ := strconv.Atoi(pidID)
	var log controller.LogRequest
	if err := validateLogBody(c, &log); err != nil {
		return
	}
	service.AddPidLog(id, log)
	c.JSON(http.StatusCreated, gin.H{"message": "Log added"})
 }

 func validatePidBody(c *gin.Context, pidModel *model.PidModel) error {
	if err := c.ShouldBindJSON(pidModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func validateLogBody(c *gin.Context, log *dto.message) error {
	if err := c.ShouldBindJSON(log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}