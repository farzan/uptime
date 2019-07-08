package domain

import (
	"math/rand"
	"time"
)

var(
	// Times in milliseconds
	minDelay = 100
	maxDelay = 15000
)

func NewHttpServiceFake() HttpServiceInterface {
	return httpServiceFake{}
}

type httpServiceFake struct {

}

func (s httpServiceFake) Process(request Request) Response {
	startTime := time.Now()
	delayInMilliseconds := minDelay + rand.Intn(maxDelay-minDelay)
	time.Sleep(time.Duration(delayInMilliseconds) * time.Millisecond)
	endTime := time.Now()

	isOk := rand.Intn(8) > 0

	if isOk {
		return Response{
			Request:    request,
			Start:      startTime,
			End:        endTime,
			StatusCode: 200,
		}
	} else {
		return Response{
			Request:    request,
			Start:      startTime,
			End:        endTime,
			StatusCode: 500,
			Error:      "Internal server error",
		}
	}
}
