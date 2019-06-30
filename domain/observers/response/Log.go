package response

import (
	"UptimeMonitor/domain/types/monitor"
	"github.com/astaxie/beego/logs"
	"time"
)

var log *logs.BeeLogger

func init() {
	log = logs.NewLogger()
	log.SetLogger(logs.AdapterFile, `{"filename": "logs/monitoring_result.log"}`)
}

type Log struct {

}

func (t Log) Notify(response monitor.Response) {
	log.Info("Timestamp: %v, Monitor Id: %v, Status: %v, Response time: %v",
		response.Start.Format(time.RFC3339),
		response.Request.Monitor.Id,
		response.StatusCode,
		response.Duration(),
	)
}