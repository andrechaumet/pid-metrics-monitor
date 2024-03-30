package display

import (
	"fmt"
	"pid-metrics-monitor/model"
	"pid-metrics-monitor/service"
	"strconv"
	"strings"
)

const (
	timeFormat   = "2006-01-02 15:04:05"
	headerFormat = "PID\tNAME\t\tLAPSED-TIME\tCURRENT / TOTAL\tPERCENTAGE %\tLAST UPDATE"
	valueFormat  = "%s\t%s\t%.2f%%\t%s\t%.2f%%\t%s\t%s\t%d\n"
)

func Display() {
	fmt.Println(headerFormat)
	processes := service.FindAll()
	for _, p := range processes {
		displayPid(p)
	}
}

func displayPid(p model.Pid) {
	id := truncateValue(strconv.Itoa(p.ID), 10)
	name := truncateValue(p.Name, 10)
	lapsedTime := p.LapsedTime()
	currentTotal := fmt.Sprintf("%d / %d", p.CurrentIterations, p.TotalIterations)
	percentage := p.Percentage()
	lastUpdate := p.LastUpdate.Format(timeFormat)
	fmt.Printf(valueFormat, id, name, lapsedTime, currentTotal, percentage, lastUpdate)
}

func truncateValue(str string, length int) string {
	if len(str) > length {
		return str[:length]
	}
	return str + strings.Repeat(" ", length-len(str))
}
