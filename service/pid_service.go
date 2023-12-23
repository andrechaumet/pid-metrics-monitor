package service

import (
	"time"
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/persistence"
)

func FindAll() []model.PidModel {
	return persistence.FindAll()
}

func Save(pid model.PidModel) {
	setLastUpdate(&pid)
	persistence.Save(pid)
}

func Update(sent model.PidModel) {
	found, exists := persistence.FindById(sent.ID)
	if exists {
		setIterations(&found, &sent)	
	}
	persistence.Save(found)
}

func setIterations(found *model.PidModel, sent *model.PidModel) {
	found.CurrentIterations = sent.CurrentIterations
	found.Logs = append(found.Logs, sent.Logs...)
}

func setLastUpdate(pid *model.PidModel) {
	pid.LastUpdate = time.Now()
}

/*func AddPidLog(pidID int, log []string) {
	pid, exists := persistence.FindById(pidID)
	if exists {
		pid.Logs = initPidLogsIfNil(pid.Logs)
		pid.Logs = append(pid.Logs, log...)
		Save(pid)
	}
}

func initPidLogsIfNil(logs []string) []string {
    if logs == nil {
        return []string{}
    }
    return logs
}*/