package services

import (
	"UptimeMonitor/domain/types/monitor"
	"math/rand"
	"time"
)

var(
	minDelay = 100
	maxDelay = 2000
)

type HttpServiceFake struct {

}

func (s HttpServiceFake) Process(request monitor.Request) monitor.Response {
	startTime := time.Now()

	delay := time.Duration(minDelay + rand.Intn(maxDelay - minDelay))
	time.Sleep(delay)

	endTime := time.Now()

	return monitor.Response{
		request,
		startTime,
		endTime,
		200,
		"",
	}
}