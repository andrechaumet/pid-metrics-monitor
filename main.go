package pid_metrics_monitor

import (
	"net/http"
	"pid-metrics-monitor/base"
)

func main() {
	r := base.SetupRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
