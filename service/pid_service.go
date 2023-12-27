package service

import (
	"time"
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/repository"
)

func FindAll() []model.PidModel {
	found := persistence.FindAll()
	for i := range found {
		calculateCurrentSpeed(&found[i])
	}
	return found
}
//TODO:
//func Create

func Save(pid model.PidModel) {
	setLastUpdate(&pid)
	persistence.Save(pid)
}

func Update(sent model.PidModel) {
	found, exists := persistence.FindById(sent.ID)
	if exists {
		metrify(&sent, &found)	
		Save(found)
	}
}

func calculateCurrentSpeed(found *model.PidModel) {
	found.CurrentSpeed = updateCurrentSpeed(0, found)
}

//I wrote the algorithm to work with seconds, but maybe it should spit millis instead
func updateCurrentSpeed(newIterations int, found *model.PidModel) float64 {
	iterationsIncrease := float64(newIterations - found.CurrentIterations)
	timeElapsedSinceLastUpdate := time.Since(found.LastUpdate).Seconds()
	currentSpeed := iterationsIncrease / timeElapsedSinceLastUpdate
	return currentSpeed
}

//TODO:
func metrify(sent, found *model.PidModel) {
	found.CurrentSpeed = updateCurrentSpeed(sent.CurrentIterations, found)
	found.CurrentIterations = sent.CurrentIterations
	found.Percentage = 
	found.LapsedTime = 
	found.ExpectedTime =
	found.Logs = append(found.Logs, sent.Logs...)
}

func setLastUpdate(pid *model.PidModel) {
	pid.LastUpdate = time.Now()
}

