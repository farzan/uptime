package monitoring

import (
	"UptimeMonitor/dataAccess/repositories"
	"UptimeMonitor/domain/types/monitor"
	workerModel "UptimeMonitor/domain/types/monitor/worker"
	"fmt"
)

func Run() {
	repository := repositories.NewAllMonitorsFake()
	monitors := repository.Get()

	for _, mon := range monitors {
		startWorker(mon)
	}
}

func startWorker(mon monitor.Monitor) {
	fmt.Printf("[%v] Starting...\n", mon.Id)
	worker := workerModel.NewWorker(mon)
	go worker.Start()
	fmt.Printf("[%v] Started\n", mon.Id)
}
