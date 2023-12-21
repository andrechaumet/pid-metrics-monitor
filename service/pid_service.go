package service

import (
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/persistence"
)

func Save(pid model.PidModel) {
	persistence.Save(pid)
}

func Update(updatedPid model.PidModel) {
	persistence.FindById(updatedPid.ID)
	persistence.Save(updatedPid)
}

func AddPidLog(pidID int, log []string) {
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
}