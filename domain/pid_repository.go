package domain

import (
	"sync"
)

type PidsMemory struct {
	sync.Mutex
	PidsMap   map[int]Pid
	currentID int
}

var pidsMemoryInstance *PidsMemory
var once sync.Once

func getPidsMemoryInstance() *PidsMemory {
	once.Do(func() {
		pidsMemoryInstance = &PidsMemory{
			PidsMap: make(map[int]Pid),
		}
	})
	return pidsMemoryInstance
}

func Create(pid Pid) int {
	pm := getPidsMemoryInstance()
	pm.Lock()
	defer pm.Unlock()
	pid.ID = pm.currentID
	pm.PidsMap[pm.currentID] = pid
	pm.currentID++
	return pid.ID
}

func FindAll() []Pid {
	pm := getPidsMemoryInstance()
	pm.Lock()
	defer pm.Unlock()
	pids := make([]Pid, 0, len(pm.PidsMap))
	for _, pid := range pm.PidsMap {
		pids = append(pids, pid)
	}
	return pids
}

func Update(pid Pid) {
	pm := getPidsMemoryInstance()
	pm.Lock()
	defer pm.Unlock()
	pm.PidsMap[pid.ID] = pid
}

func FindById(ID int) (*Pid, bool) {
	pm := getPidsMemoryInstance()
	pm.Lock()
	defer pm.Unlock()
	pid, exists := pm.PidsMap[ID]
	return &pid, exists
}
