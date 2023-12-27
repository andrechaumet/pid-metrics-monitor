package batch

import (
	"fmt"
	"time"

	"pid-metrics-monitor/service"
)

func LoadAndDisplay() {
	for {
		display()
		time.Sleep(time.Second)
	}
}

func display() {
	for _, pid := range service.FindAll() {
		fmt.Printf("ID: [%d], StartTime: [%s], LastUpdate: [%s]\n", pid.ID, pid.StartTime.Format(time.RFC3339), pid.LastUpdate.Format(time.RFC3339))
		fmt.Printf("Iterations: [%d/%d], Percentage: %.2f%%, Speed: %.2f iter/s\n", pid.CurrentIterations, pid.TotalIterations, pid.Percentage, pid.CurrentSpeed)
		fmt.Printf("LapsedTime: [%d] seconds, ExpectedTime: [%s]\n", pid.LapsedTime, pid.ExpectedTime.Format(time.RFC3339))
	}
}