package domain

func Run() {
	// todo use a service provider
	monitors := GetMonitorService().FindAll()

	for _, mon := range monitors {
		StartWorker(&mon)
	}
}

func StartWorker(mon *Monitor) {
	//debug.Printf("[%v] Starting...\n", mon.Id)
	worker := createWorker(*mon)
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
