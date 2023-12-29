package monitor

import (
	"pid-metrics-monitor/service"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func Display() {
	for {
		display()
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		time.Sleep(time.Second)
	}
}

func display() {
	for {
		//clearLastLines(3)
		time.Sleep(time.Second)
		for _, pid := range service.FindAll() {
			fmt.Printf("ID: %d, StartTime: %s, LastUpdate: %s\n", pid.ID, pid.StartTime.Format(timeFormat), pid.LastUpdate.Format(timeFormat))
			fmt.Printf("Iterations: %d/%d, Percentage: %.2f%%, Speed: %.2f iter/s\n", pid.CurrentIterations, pid.TotalIterations, pid.Percentage, pid.CurrentSpeed)
			fmt.Printf("LapsedTime: %d seconds, ExpectedFinishTime: %s\n", pid.LapsedTime, pid.ExpectedTime.Format(timeFormat))
			fmt.Printf("--")
		}
	}
}

func clearLastLines(numLines int) {
	for i := 0; i < numLines; i++ {
		fmt.Print("\033[A\033[K")
	}
}