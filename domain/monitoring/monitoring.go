package monitoring

import (
	"UptimeMonitor/dataAccess/repositories"
	"UptimeMonitor/domain/types/monitor"
)

func Run() {
	repository := repositories.NewAllMonitorsFake()
	monitors := repository.Get()

	for _, mon := range monitors {
		startWorker(mon)
	}
}

func startWorker(mon monitor.Monitor) {
	//fmt.Println(monitor)
	worker := monitor.NewWorker(mon)
	worker.Start()
}
