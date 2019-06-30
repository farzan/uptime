package worker

import (
	domainContracts "UptimeMonitor/domain/contracts/domain"
	"UptimeMonitor/domain/services"
	"UptimeMonitor/domain/types/monitor"
	"fmt"
	"math"
	"time"
)

type Worker struct {
	monitor    monitor.Monitor
	ResponseCh chan monitor.Response
	stopCh     chan bool
}

func NewWorker(mon monitor.Monitor) Worker {
	return Worker{
		mon,
		make(chan monitor.Response),
		make(chan bool),
	}
}

func (w Worker) Start() {
	for {
		go sendRequest(w)

		response := waitForResponse(w)

		delay := getDelay(response)
		fmt.Printf("[%v] Wait %v\n", w.monitor.Id, delay)
		time.Sleep(delay)
	}
}

func getDelay(response monitor.Response) time.Duration {
	requestDuration := response.End.Sub(response.Start)
	delay := time.Duration(response.Request.Monitor.Interval) * time.Second - requestDuration
	normalizedDelay := time.Duration(math.Max(0, float64(delay)))

	return normalizedDelay
}

func sendRequest(worker Worker) {
	request := monitor.Request{Monitor: worker.monitor}

	worker.ResponseCh <- getResponse(request)
}

func waitForResponse(worker Worker) monitor.Response {
	for {
		select {
		case response := <- worker.ResponseCh:
			fmt.Printf("[%v] Response received; Dur: %v\n",
				response.Request.Monitor.Id,
				response.End.Sub(response.Start),
			)
			return response
		default: <- time.After(500 * time.Millisecond)
			fmt.Print(".")
		}
	}
}

func getResponse(request monitor.Request) monitor.Response {
	return getHttpService().Process(request)
}

func getHttpService() domainContracts.HttpService {
	return services.HttpServiceFake{}
}

func (w Worker) Stop() {

}
