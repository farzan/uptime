package monitor

import (
	domainContracts "UptimeMonitor/domain/contracts/domain"
	"UptimeMonitor/domain/services"
	"time"
)

type Worker struct {
	monitor Monitor
	ResponseCh chan Response
	stopCh chan bool
}

func NewWorker(monitor Monitor) Worker {
	return Worker{
		monitor,
		make(chan Response),
		make(chan bool),
	}
}

func (w Worker) Start() {
	for {
		go sendRequest(w)

		select {
		case response := <- w.ResponseCh:

		default:

		}

		time.Sleep(5 * time.Second)
	}
}

func sendRequest(worker Worker) {
	request := Request{worker.monitor}

	worker.ResponseCh <- getResponse(request)
}

func getResponse(request Request) Response {
	return getHttpService().Process(request)
}

func getHttpService() domainContracts.HttpService {
	return services.HttpServiceFake{}
}

func (w Worker) Stop() {

}
