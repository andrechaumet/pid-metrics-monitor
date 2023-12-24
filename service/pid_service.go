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
		setIterations(&found, &sent)	
	}
	persistence.Save(found)
}


func setIterations(found *model.PidModel, sent *model.PidModel) {
	found.CurrentIterations = sent.CurrentIterations
	found.Logs = append(found.Logs, sent.Logs...)
}

func setLastUpdate(pid *model.PidModel) {
	pid.LastUpdate = time.Now()
}

func calculateCurrentSpeed(sent model.PidModel, found model.PidModel) float64 {
	timeElapsed := time.Now().Sub(found.StartTime).Seconds()
	iterationsIncrease := sent.CurrentIterations - found.CurrentIterations
	timeElapsedSinceLastUpdate := timeElapsed.sub(found.LastUpdate).Seconds()
	currentSpeed := float64(iterationsIncrease) / timeElapsedSinceLastUpdate
	return currentSpeed
}

func metrify(found *model.PidModel, sent *model.PidModel) {
	found.LastUpdate = time.Now()
	found.CurrentIterations = sent.CurrentIterations
	found.CurrentSpeed = calculateCurrentSpeed(found)

}

