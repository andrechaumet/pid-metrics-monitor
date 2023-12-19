package persistence

import "pid-metrics-monitor/model"

type PidsMemory struct {
	PidsMap map[int]model.PidModel
}

func NewPidsMemory() *PidsMemory {
	return &PidsMemory{
		PidsMap: make(map[int]model.PidModel),
	}
}

func (p *PidsMemory) FindById(ID int) (model.PidModel, bool) {
    pid, exists := p.PidsMap[ID]
    if exists {
        return pid
    }
    return exists
}

func (p *PidsMemory) Save(pid model.PidModel) {
	p.PidsMap[pid.ID] = pid
}