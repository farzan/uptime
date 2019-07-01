package services

import (
	"UptimeMonitor/domain/types/monitor"
	debug "UptimeMonitor/utils"
	"math/rand"
	"time"
)

var(
	// Times in milliseconds
	minDelay = 100
	maxDelay = 15000
)

type HttpServiceFake struct {

}

func (s HttpServiceFake) Process(request monitor.Request) monitor.Response {
	startTime := time.Now()

	delayInMilliseconds := minDelay + rand.Intn(maxDelay-minDelay)
	debug.Printf("[%v] Request will take %v ms\n", request.Monitor.Id, delayInMilliseconds)
	time.Sleep(time.Duration(delayInMilliseconds) * time.Millisecond)

	endTime := time.Now()

	return monitor.Response{
		request,
		startTime,
		endTime,
		200,
		"",
	}
}
