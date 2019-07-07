package domain

type WorkerServiceInterface interface {
	Start(monitor *Monitor)
	Stop(monitorId int)
}

func NewWorkerService() WorkerServiceInterface {
	w := workerService{}
	w.workers = make(map[int]Worker)

	return &w
}

type workerService struct {
	workers map[int]Worker
}

func (s *workerService) Start(monitor *Monitor) {
	//debug.Printf("[%v] Starting...\n", mon.Id)
	worker := s.createWorker(*monitor)
	s.workers[monitor.Id] = worker

	go worker.Start()
	//debug.Printf("[%v] Started\n", mon.Id)
}

func (s *workerService) Stop(monitorId int) {
	s.workers[monitorId].Stop()
	delete(s.workers, monitorId)
}

func (s workerService) createWorker(mon Monitor) Worker {
	worker := NewWorker(mon)

	worker.AttachResponseObserver(TerminalObserver{})
	worker.AttachResponseObserver(LogObserver{})
	worker.AttachResponseObserver(DatabaseObserver{})

	return worker
}