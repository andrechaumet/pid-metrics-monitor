package monitor

import (
	"pid-metrics-monitor/service"
	"fmt"
	"time"
	"os"
	"os/exec"
	"runtime"
)

const timeFormat = "2006-01-02 15:04:05"

func Display() {
	for {
		display()
	}
}

func display() {
	for {
		clear()
		time.Sleep(time.Second)
		for _, pid := range service.FindAll() {
			fmt.Printf("ID: %d, StartTime: %s, LastUpdate: %s\n", pid.ID, pid.StartTime.Format(timeFormat), pid.LastUpdate.Format(timeFormat))
			fmt.Printf("Iterations: %d/%d, Percentage: %.2f%%, Speed: %.2f iter/s\n", pid.CurrentIterations, pid.TotalIterations, pid.Percentage, pid.CurrentSpeed)
			fmt.Printf("LapsedTime: %d seconds, ExpectedFinishTime: %s\n", pid.LapsedTime, pid.ExpectedTime.Format(timeFormat))
			fmt.Printf("--")
		}
	}
}

func clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("No se puede limpiar la consola en este sistema operativo.")
	}
}