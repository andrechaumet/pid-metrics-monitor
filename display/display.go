package display

import (
	"fmt"
	"os"
	"os/exec"
	"pid-metrics-monitor/service"
	"runtime"
	"strings"
	"time"
)

const timeFormat = "2006-11-02 15:04:05"
const idDisplaySize = 10
const nameDisplaySize = 10
const timeDisplaySize = 10
const percentageDisplaySize = 10

func Display() {
	for {
		time.Sleep(3 * time.Second)
		for _, pid := range service.FindAll() {
			fmt.Printf("ID      NAME      LAPSED TIME      CURRENT / TOTAL      PERCENTAGE      LAST UPDATE")
			fmt.Printf("%s %s %d %d/%d %.2f%% %s",
				formatValue(pid.ID, idDisplaySize),
				formatValue(pid.Name, nameDisplaySize),
				formatValue(pid.LapsedTime(), timeDisplaySize),
				formatValue(pid.Percentage(), percentageDisplaySize),
				formatValue(pid.LastUpdate, timeDisplaySize),
			)
			fmt.Println("--")
		}
	}
}

//ID, NAME, LAPSED-TIME, CURRENT/TOTAL, PERCENTAGE, LAST UPDATE

func formatValue(value interface{}, length int) string {
	str := fmt.Sprintf("%v", value)
	if len(str) > length {
		return str[:length]
	} else if len(str) < length {
		spaces := length - len(str)
		return str + strings.Repeat(" ", spaces)
	}
	return str
}

func formatLapsedTime() {

}

func formatCurrentTotal() {

}

func formatPercentage() {

}

func formatLastUpdate() {

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
