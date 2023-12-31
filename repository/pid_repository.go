package persistence

import "pid-metrics-monitor/model"

type PidsMemory struct {
	PidsMap map[int]model.PidModel
}

var pidsMemoryInstance = newPidsMemory()

func newPidsMemory() *PidsMemory {
	return &PidsMemory{
		PidsMap: make(map[int]model.PidModel),
	}
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

func FindById(ID int) (model.PidModel, bool) {
	pid, exists := pidsMemoryInstance.PidsMap[ID]
	return pid, exists
}

func (pm *PidsMemory) findById(ID int) model.PidModel {
	return pm.PidsMap[ID]
}