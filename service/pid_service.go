package service

import (
	"time"
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/persistence"
)

func Save(pid model.PidModel) {
	persistence.Save(pid)
}

func Update(sent model.PidModel) {
	found, exists := persistence.FindById(sent.ID)
	if exists {
		updateFound(&found, &sent)	
	}
	persistence.Save(found)
}

func updateFound(found *model.PidModel, sent *model.PidModel) {
	found.LastUpdate = time.Now()
	found.CurrentIterations = sent.CurrentIterations
	found.Logs = append(found.Logs, sent.Logs...)
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