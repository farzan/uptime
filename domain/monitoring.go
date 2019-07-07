package domain

func Run() {
	// todo use a service provider
	monitors := GetMonitorService().FindAll()

	for _, mon := range monitors {
		GetWorkerService().Start(&mon)
	}
}
