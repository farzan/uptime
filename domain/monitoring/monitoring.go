package monitoring

import (
	"UptimeMonitor/dataAccess/repositories"
	"UptimeMonitor/domain/observers/response"
	"UptimeMonitor/domain/types/monitor"
	workerModel "UptimeMonitor/domain/types/monitor/worker"
)

func Run() {
	repository := repositories.NewAllMonitorsRepositoryFake()
	monitors := repository.All()

	for _, mon := range monitors {
		startWorker(mon)
	}
}

func startWorker(mon monitor.Monitor) {
	//debug.Printf("[%v] Starting...\n", mon.Id)
	worker := createWorker(mon)
	go worker.Start()
	//debug.Printf("[%v] Started\n", mon.Id)
}

func createWorker(mon monitor.Monitor) workerModel.Worker {
	worker := workerModel.NewWorker(mon)

	//worker.AttachResponseObserver(response.Terminal{})
	worker.AttachResponseObserver(response.Log{})
	worker.AttachResponseObserver(response.Database{})

	return worker
}
