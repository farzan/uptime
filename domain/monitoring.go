package domain

func Run() {
	monitors := GetMonitorService().FindAll()

	for _, mon := range monitors {
		GetWorkerService().Start(&mon)
	}
}
