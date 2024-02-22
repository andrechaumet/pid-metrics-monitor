package persistence

import (
	"pid-metrics-monitor/model"
)

type PidsMemory struct {
	PidsMap   map[int]model.PidModel
	currentID int
}

var pidsMemoryInstance = newPidsMemory()

func newPidsMemory() *PidsMemory {
	return &PidsMemory{
		PidsMap:   make(map[int]model.PidModel),
		currentID: 0,
	}
}

func Create(pid model.PidModel) int {
	return pidsMemoryInstance.create(pid)
}

func (pm *PidsMemory) create(pid model.PidModel) int {
	pid.ID = pm.currentID
	pm.PidsMap[pm.currentID] = pid
	pm.currentID++
	return pid.ID
}

func FindAll() []model.PidModel {
	pids := make([]model.PidModel, 0, len(pidsMemoryInstance.PidsMap))
	for _, pid := range pidsMemoryInstance.PidsMap {
		pids = append(pids, pid)
	}
	return pids
}

func Save(pid model.PidModel) {
	pidsMemoryInstance.save(pid)
}

func (pm *PidsMemory) save(pid model.PidModel) {
	pm.PidsMap[pid.ID] = pid
}

func FindById(ID int) model.PidModel {
	pid := pidsMemoryInstance.PidsMap[ID]
	return pid
}

func (pm *PidsMemory) findById(ID int) model.PidModel {
	return pm.PidsMap[ID]
}
