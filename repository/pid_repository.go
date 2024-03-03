package persistence

import (
	"pid-metrics-monitor/model"
)

type PidsMemory struct {
	PidsMap   map[int]model.Pid
	currentID int
}

var pidsMemoryInstance = newPidsMemory()

func newPidsMemory() *PidsMemory {
	return &PidsMemory{
		PidsMap:   make(map[int]model.Pid),
		currentID: 0,
	}
}

func Create(pid model.Pid) int {
	return pidsMemoryInstance.create(pid)
}

func (pm *PidsMemory) create(pid model.Pid) int {
	pid.ID = pm.currentID
	pm.PidsMap[pm.currentID] = pid
	pm.currentID++
	return pid.ID
}

func FindAll() []model.Pid {
	pids := make([]model.Pid, 0, len(pidsMemoryInstance.PidsMap))
	for _, pid := range pidsMemoryInstance.PidsMap {
		pids = append(pids, pid)
	}
	return pids
}

func Save(pid model.Pid) {
	pidsMemoryInstance.save(pid)
}

func (pm *PidsMemory) save(pid model.Pid) {
	pm.PidsMap[pid.ID] = pid
}

func FindById(ID int) *model.Pid {
	pid, exists := pidsMemoryInstance.PidsMap[ID]
	if !exists {
		return nil
	}
	return &pid
}

func (pm *PidsMemory) findById(ID int) model.Pid {
	return pm.PidsMap[ID]
}
