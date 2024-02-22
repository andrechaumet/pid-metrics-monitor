package service

import (
	model "pid-metrics-monitor/model"
	persistence "pid-metrics-monitor/repository"
	"time"
)

func FindAll() []model.PidModel {
	found := persistence.FindAll()
	return found
}

func Create(pid model.PidModel) {
	setStartTime(&pid)
	setLastUpdate(&pid)
	persistence.Save(pid)
}

func Update(sentPid model.PidModel) {
	foundPid, exists := persistence.FindById(sentPid.ID)
	if exists {
		//metrify(&sentPid, &foundPid)
		save(foundPid)
	}
}

func save(pid model.PidModel) {
	setLastUpdate(&pid)
	persistence.Save(pid)
}

func setStartTime(pid *model.PidModel) {
	if pid.StartTime == nil {
		pid.StartTime = new(time.Time)
	}
	*pid.StartTime = time.Now()
}

func setLastUpdate(pid *model.PidModel) {
	if pid.LastUpdate == nil {
		pid.LastUpdate = new(time.Time)
	}
	*pid.LastUpdate = time.Now()
}

/*


func calculateExpectedTime(sent, found model.PidModel) time.Time {
	remainingItems := float64(found.TotalIterations - sent.CurrentIterations)
	remainingSeconds := remainingItems / found.CurrentSpeed
	expectedTime := time.Now().Add(time.Duration(int64(remainingSeconds)) * time.Second)
	return expectedTime
}

// TODO: cs/s =  (current amount of iterations - previous amount of iterations) / current time - last update in seconds
func updateCurrentSpeed(newIterations int, found *model.PidModel) float64 {
	iterationsIncrease := float64(newIterations - found.CurrentIterations)
	timeElapsedSinceLastUpdate := time.Since(*found.LastUpdate).Seconds()
	if timeElapsedSinceLastUpdate == 0 {
		return 0.0
	}
	currentSpeed := iterationsIncrease / timeElapsedSinceLastUpdate
	return currentSpeed
}

// TODO: Lapsed time should be updated in Calculate curent speed
func metrify(sent, found *model.PidModel) {
	found.CurrentSpeed = updateCurrentSpeed(sent.CurrentIterations, found)
	found.CurrentIterations = sent.CurrentIterations
	found.Percentage = float64((sent.CurrentIterations * 100) / found.TotalIterations)
	found.ExpectedTime = calculateExpectedTime(*sent, *found)
	found.Logs = append(found.Logs, sent.Logs...)
}

*/
