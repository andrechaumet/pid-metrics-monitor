package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pid-metrics-monitor/service"
	"pid-metrics-monitor/model"
 )
 
 func CreatePID(c *gin.Context) {
	var pidModel model.PidModel
	if err := validateBody(c, &pidModel); err != nil {
		return
	}
	service.Save(pidModel)
	c.JSON(http.StatusCreated, gin.H{"message": "PID status created"})
}
 
 func UpdatePID(c *gin.Context) {
	pidID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "PID status updated", "id": pidID})
 }
 
 func CreatePIDLog(c *gin.Context) {
	pidID := c.Param("id")
	c.JSON(http.StatusCreated, gin.H{"message": "Log inserted", "id": pidID})
 }

 func validateBody(c *gin.Context, pidModel *model.PidModel) error {
	if err := c.ShouldBindJSON(pidModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}