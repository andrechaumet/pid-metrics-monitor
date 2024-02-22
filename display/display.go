package display

import (
	"fmt"
	"os"
	"os/exec"
	"pid-metrics-monitor/service"
	"runtime"
	"time"
)

const timeFormat = "0001-01-01 01:01:01"

func Display() {
	for {
		display()
	}
}

func display() {
	for {
		/*clear()*/
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		for _, pid := range service.FindAll() {

			fmt.Printf("%s\n", pid.Name)
			fmt.Printf("StartTime: %s, LastUpdate: %s\n", (*pid.StartTime).String(), (*pid.LastUpdate).String())
			fmt.Printf("Iterations: %d/%d\n", pid.CurrentIterations, pid.TotalIterations)
			fmt.Printf("Percentage: %.2f%%\n", pid.Percentage)
			fmt.Printf("Speed: %.2f iter/s\n", pid.CurrentSpeed())
			fmt.Printf("LapsedTime : %d", pid.LapsedTime())

			/*			fmt.Printf("%s\n", pid.Name)

						fmt.Printf("ID: %d, StartTime: %s, LastUpdate: %s\n", pid.ID, pid.StartTime.Format(timeFormat), pid.LastUpdate.Format(timeFormat))

						fmt.Printf("Iterations: %d/%d, Percentage: %.2f%%, Speed: %.2f iter/s\n", pid.CurrentIterations, pid.TotalIterations, pid.Percentage(), pid.CurrentSpeed())

						fmt.Printf("LapsedTime: %d seconds, ExpectedFinishTime: %s\n", pid.LapsedTime(), time.Now().Add(time.Duration(pid.ExpectedTime())*time.Second).Format(time.RFC3339))
			*/
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
