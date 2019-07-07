package domain

import (
	"math"
	"time"
)

type Worker interface {
	Start()
	Stop()
	AttachResponseObserver(observer ResponseObserver)
}

type worker struct {
	monitor       Monitor
	ResponseCh    chan Response
	stopCh        chan bool
	respObservers []ResponseObserver
}

func NewWorker(mon Monitor) Worker {
	return &worker{
		monitor:    mon,
		ResponseCh: make(chan Response),
		stopCh:     make(chan bool),
	}
}

func (w worker) Start() {
	for {
		go sendRequest(w)

		response := waitForResponse(w)

		select {
		case <- w.stopCh:
			return
		default:
		}

		delay := getDelay(response)
		//debug.Printf("[%v] Wait %v\n", w.monitor.Id, delay)
		time.Sleep(delay)
	}
}

func getDelay(response Response) time.Duration {
	requestDuration := response.Duration()
	delay := time.Duration(response.Request.Monitor.Interval) * time.Second - requestDuration
	normalizedDelay := time.Duration(math.Max(0, float64(delay)))

	return normalizedDelay
}

func sendRequest(worker worker) {
	request := Request{Monitor: worker.monitor}

	worker.ResponseCh <- getResponse(request)
}

func waitForResponse(worker worker) Response {
	for {
		select {
		case response := <- worker.ResponseCh:
			worker.notifyResponseObservers(response)
			return response
		//default: <- time.After(500 * time.Millisecond)
		//	debug.Print(".")
		}
	}
}

func getResponse(request Request) Response {
	return getHttpService().Process(request)
}

func getHttpService() HttpServiceInterface {
	return HttpServiceFake{}
}

func (w worker) Stop() {
	w.stopCh <- true
}

func (w *worker) AttachResponseObserver(observer ResponseObserver) {
	w.respObservers = append(w.respObservers, observer)
}

func (w worker) notifyResponseObservers(response Response) {
	for _, observer := range w.respObservers {
		observer.Notify(response)
	}
}