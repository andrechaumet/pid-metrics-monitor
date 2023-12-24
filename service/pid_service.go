package service

import (
	"time"
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/persistence"
)

func FindAll() []model.PidModel {
	return persistence.FindAll()
}

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

func calculateCurrentSpeed(sent, found *model.PidModel) float64 {
	iterationsIncrease := float64(sent.CurrentIterations - found.CurrentIterations)
	timeElapsedSinceLastUpdate := time.Since(found.LastUpdate).Seconds()
	currentSpeed := iterationsIncrease / timeElapsedSinceLastUpdate
	return currentSpeed
}

func metrify(sent, found *model.PidModel) {
	found.CurrentSpeed = calculateCurrentSpeed(sent, found)
	found.CurrentIterations = sent.CurrentIterations
	found.Logs = append(found.Logs, sent.Logs...)
}

func setLastUpdate(pid *model.PidModel) {
	pid.LastUpdate = time.Now()
}

