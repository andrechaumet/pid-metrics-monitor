package service

import "pid-metrics-monitor/persistence"
import "pid-metrics-monitor/model"

func () Save(pid model.PidModel) {
	persistence.Save(pid)
}

func Update(updatedPid model.PidModel) {
	if _, exists := p.PidsMap[updatedPid.ID]; exists {
		p.PidsMap[updatedPid.ID] = updatedPid
	}
}

func AddPidLog(ID int, log string) {
    pid, exists := persistence.FindById(ID)
    if exists {
        initLogsIfNil(&pid.Logs)
        pid.Logs = append(pid.Logs, log)
        persistence.Save(&pid)
    }
}

func initLogsIfNil(logs *[]string) {
    if *logs == nil {
        *logs = []string{}
    }
}

