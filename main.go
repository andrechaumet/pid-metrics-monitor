package pid_metrics_monitor

import (
	"net/http"
	"pid-metrics-monitor/base"
)

func main() {
	r := base.SetupRouter()
	http.ListenAndServe(":8081", r)
}
