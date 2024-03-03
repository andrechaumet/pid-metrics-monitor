package service

import (
	"pid-metrics-monitor/model"
	persistence "pid-metrics-monitor/repository"
	"time"
)

func FindAll() []model.Pid {
	found := persistence.FindAll()
	return found
}

func Create(pid *model.Pid) int {
	updateTime(&pid.StartTime)
	updateTime(&pid.LastUpdate)
	return persistence.Create(pid)
}

func Update(sentPid model.Pid) {
	foundPid := persistence.FindById(sentPid.ID)
	if foundPid != nil {
		//TODO: complete later
		foundPid.CurrentIterations = sentPid.CurrentIterations
		save(*foundPid)
	}
}

func save(pid *model.Pid) {
	persistence.Save(pid)
}

func updateTime(t *time.Time) {
	*t = time.Now()
}
