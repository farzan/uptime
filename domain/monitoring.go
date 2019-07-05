package domain

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

	worker.AttachResponseObserver(TerminalObserver{})
	worker.AttachResponseObserver(LogObserver{})
	worker.AttachResponseObserver(DatabaseObserver{})

	return worker
}
