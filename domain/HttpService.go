package domain

import (
	"net/http"
	"time"
)

func NewHttpService() HttpServiceInterface {
	return httpService{}
}

type httpService struct {

}

func (s httpService) Process(request Request) Response {
	startTime := time.Now()
	_, err := http.Get(request.Monitor.Url)
	endTime := time.Now()

	if err == nil {
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
			Error:      err.Error(),
		}
	}
}
