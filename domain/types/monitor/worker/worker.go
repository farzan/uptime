package worker

import (
	domainContracts "UptimeMonitor/domain/contracts/services"
	"UptimeMonitor/domain/services"
	"UptimeMonitor/domain/types/monitor"
	debug "UptimeMonitor/utils"
	"math"
	"time"
)

type Worker interface {
	Start()
	Stop()
	AttachResponseObserver(observer domainContracts.ResponseObserver)
}

type worker struct {
	monitor    		monitor.Monitor
	ResponseCh 		chan monitor.Response
	stopCh     		chan bool
	respObservers	[]domainContracts.ResponseObserver
}

func NewWorker(mon monitor.Monitor) Worker {
	return &worker{
		monitor:    mon,
		ResponseCh: make(chan monitor.Response),
		stopCh:     make(chan bool),
	}
}

func (w worker) Start() {
	for {
		go sendRequest(w)

		response := waitForResponse(w)

		delay := getDelay(response)
		debug.Printf("[%v] Wait %v\n", w.monitor.Id, delay)
		time.Sleep(delay)
	}
}

func getDelay(response monitor.Response) time.Duration {
	requestDuration := response.Duration()
	delay := time.Duration(response.Request.Monitor.Interval) * time.Second - requestDuration
	normalizedDelay := time.Duration(math.Max(0, float64(delay)))

	return normalizedDelay
}

func sendRequest(worker worker) {
	request := monitor.Request{Monitor: worker.monitor}

	worker.ResponseCh <- getResponse(request)
}

func waitForResponse(worker worker) monitor.Response {
	for {
		select {
		case response := <- worker.ResponseCh:
			worker.notifyResponseObservers(response)
			return response
		default: <- time.After(500 * time.Millisecond)
			debug.Print(".")
		}
	}
}

func getResponse(request monitor.Request) monitor.Response {
	return getHttpService().Process(request)
}

func getHttpService() domainContracts.HttpService {
	return services.HttpServiceFake{}
}

func (w worker) Stop() {

}

func (w *worker) AttachResponseObserver(observer domainContracts.ResponseObserver) {
	w.respObservers = append(w.respObservers, observer)
}

func (w worker) notifyResponseObservers(response monitor.Response) {
	for _, observer := range w.respObservers {
		observer.Notify(response)
	}
}