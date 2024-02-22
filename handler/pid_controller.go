package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func Create(c *gin.Context) {
	var pid PidDto
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	service.Create(toModel(pid))
	c.Status(http.StatusCreated)
}

/*func save(c *gin.Context) {
	var pid PidDto
	if err := bindAndValidate(c, &pid); err != nil {
		return
	}
	service.save(toModel(pid))
	c.Status(http.StatusCreated)
}
*/

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
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	return nil
}

func FindAll(c *gin.Context) {
	pids := service.FindAll()
	c.JSON(http.StatusOK, pids)
}

func toModel(dto PidDto) model.PidModel {
	var domain model.PidModel
	domain.CurrentIterations = dto.CurrentIterations
	domain.TotalIterations = dto.TotalIterations
	domain.Name = dto.Name
	domain.Logs = dto.Logs
	return domain
}
