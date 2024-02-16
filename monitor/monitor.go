package monitor

import (
	"pid-metrics-monitor/service"
	"fmt"
	"time"
	"os"
	"os/exec"
	"runtime"
)

const timeFormat = "0001-01-01 01:01:01"

func Display() {
	for {
		display()
	}
}

func display() {
	for {
		//clear()
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		for _, pid := range service.FindAll() {
			fmt.Printf("ID: %d, StartTime: %s, LastUpdate: %s\n", pid.ID, pid.StartTime.Format(timeFormat), pid.LastUpdate.Format(timeFormat))
			fmt.Printf("Iterations: %d/%d, Percentage: %.2f%%, Speed: %.2f iter/s\n", pid.CurrentIterations, pid.TotalIterations, pid.Percentage, pid.CurrentSpeed)
			fmt.Printf("LapsedTime: %d seconds, ExpectedFinishTime: %s\n", pid.LapsedTime, pid.ExpectedTime.Format(timeFormat))
			fmt.Printf("--")
			fmt.Println("--")
		}
	}
}

func clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}