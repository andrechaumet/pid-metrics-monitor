package persistence

import "pid-metrics-monitor/model"

type PidsMemory struct {
	PidsMap map[int]model.PidModel
}

var pidsMemoryInstance = NewPidsMemory()

func NewPidsMemory() *PidsMemory {
	return &PidsMemory{
		PidsMap: make(map[int]model.PidModel),
	}
}

func Save(pid model.PidModel) {
	pidsMemoryInstance.save(pid)
}

func (pm *PidsMemory) save(pid model.PidModel) {
	pm.PidsMap[pid.ID] = pid
}

func FindById(ID int) (model.PidModel, bool) {
	pid, exists := pidsMemoryInstance.PidsMap[ID]
	return pid, exists
}

func (pm *PidsMemory) findById(ID int) model.PidModel {
	return pm.PidsMap[ID]
}