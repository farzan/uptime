package monitoring

import (
	"UptimeMonitor/dataAccess/repositories"
	"UptimeMonitor/domain/observers/response"
	"UptimeMonitor/domain/types/monitor"
	workerModel "UptimeMonitor/domain/types/monitor/worker"
	debug "UptimeMonitor/utils"
)

func Run() {
	repository := repositories.NewAllMonitorsFake()
	monitors := repository.Get()

	for _, mon := range monitors {
		startWorker(mon)
	}
}

func startWorker(mon monitor.Monitor) {
	debug.Printf("[%v] Starting...\n", mon.Id)
	worker := createWorker(mon)
	go worker.Start()
	debug.Printf("[%v] Started\n", mon.Id)
}

func createWorker(mon monitor.Monitor) workerModel.Worker {
	worker := workerModel.NewWorker(mon)

	worker.AttachResponseObserver(response.Terminal{})
	worker.AttachResponseObserver(response.Log{})
	worker.AttachResponseObserver(response.Database{})

	return worker
}
