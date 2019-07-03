package domain

import (
	debug "UptimeMonitor/utils"
)

func NewMonitoringResultsServiceFake() MonitoringResultServiceInterface {
	s := monitoringResultServiceFake{}
	s.responses = make(map[int] []Response)

	return s
}

type monitoringResultServiceFake struct {
	responses map[int] []Response
}

func (s monitoringResultServiceFake) StoreMonitoringResponse(response Response) {
	responseEntity := Response{
		MonitorId: response.Request.Monitor.Id,
		Start: response.Start,
		End: response.End,
		StatusCode: response.StatusCode,
		Error: response.Error,
	}

	s.StoreResponse(responseEntity)
}

func (s monitoringResultServiceFake) StoreResponse(response Response) {
	s.responses[response.MonitorId] = append([]Response{response}, s.responses[response.MonitorId]...)

	// Debug code:
	lengths := make(map[int]int)
	for k, v := range s.responses {
		lengths[k] = len(v)
	}
	debug.Printf("=====response stored; count: %v\n", lengths)
}

func (s monitoringResultServiceFake) FindResponseById(id int) (Response, error) {
	// todo
	return Response{}, errorString("")
}

func (s monitoringResultServiceFake) FindResponsesByFilter(filter ResponseFilter) []Response {
	monId := filter.GetMonitorId()
	set := s.responses[monId]
	pos := len(set) - filter.GetCount()

	if pos < 0 {
		pos = 0
	}

	return set[pos:]
}

type errorString string

func (e errorString) Error() string {
	return string(e)
}