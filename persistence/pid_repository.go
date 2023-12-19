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
    return pid, exists
}

func (p *PidsMemory) Save(pid model.PidModel) {
	pid.Logs = []string{}
	p.PidsMap[pid.ID] = pid
}