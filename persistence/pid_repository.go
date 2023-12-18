package persistence

import "example/user/hello/model"

type PidsMemory struct {
	PidsMap map[int]model.PidModel
}

func NewPidsMemory() *PidsMemory {
	return &PidsMemory{
		PidsMap: make(map[int]model.PidModel),
	}
}

func (p *PidsMemory) Save(pid model.PidModel) {
	p.PidsMap[pid.ID] = pid
}

func (p *PidsMemory) Update(updatedPid model.PidModel) {
	if _, exists := p.PidsMap[updatedPid.ID]; exists {
		p.PidsMap[updatedPid.ID] = updatedPid
	}
}

func (p *PidsMemory) AddPidLog(ID int, newLog string) {
	if pid, exists := p.PidsMap[ID]; exists {
		if pid.Logs == nil {
			pid.Logs = []string{}
		}
		pid.Logs = append(pid.Logs, newLog)
		p.PidsMap[ID] = pid
	}
}