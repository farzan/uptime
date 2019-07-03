package domain

import (
	"UptimeMonitor/domain/observers/response"
)

func Run() {
	// todo use a service provider
	monitors := GetMonitorService().GetAll()

	for _, mon := range monitors {
		startWorker(mon)
	}
}

func startWorker(mon Monitor) {
	//debug.Printf("[%v] Starting...\n", mon.Id)
	worker := createWorker(mon)
	go worker.Start()
	//debug.Printf("[%v] Started\n", mon.Id)
}

func createWorker(mon Monitor) Worker {
	worker := NewWorker(mon)

	//worker.AttachResponseObserver(response.Terminal{})
	worker.AttachResponseObserver(response.Log{})
	worker.AttachResponseObserver(response.Database{})

	return worker
}
