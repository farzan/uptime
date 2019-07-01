package services

import (
	"UptimeMonitor/domain/contracts/services"
	"UptimeMonitor/domain/entities"
	"UptimeMonitor/domain/services/filters"
	"UptimeMonitor/domain/types/monitor"
	"fmt"
)

func NewMonitoringResultsServiceFake() services.MonitoringResultService {
	s := monitoringResultServiceFake{}
	s.responses = make(map[int] []entities.Response)

	return s
}

type monitoringResultServiceFake struct {
	responses map[int] []entities.Response
}

func (s monitoringResultServiceFake) StoreMonitoringResponse(response monitor.Response) {
	responseEntity := entities.Response{
		MonitorId: response.Request.Monitor.Id,
		Start: response.Start,
		End: response.End,
		StatusCode: response.StatusCode,
		Error: response.Error,
	}

	s.StoreResponse(responseEntity)
}

func (s monitoringResultServiceFake) StoreResponse(response entities.Response) {
	s.responses[response.MonitorId] = append(s.responses[response.MonitorId], response)

	lengths := make(map[int]int)
	for k, v := range s.responses {
		lengths[k] = len(v)
	}

	fmt.Printf("=====response stored; count: %v\n", lengths)
}

func (s monitoringResultServiceFake) FindResponseById(id int) (entities.Response, error) {
	// todo
	return entities.Response{}, errorString("")
}

func (s monitoringResultServiceFake) FindResponsesByFilter(filter filters.ResponseFilter) []entities.Response {
	monId := filter.GetMonitorId()
	set := s.responses[monId]
	pos := len(set) - filter.GetCount()

	fmt.Printf("===============Pos: %v", pos)

	if pos < 0 {
		pos = 0
	}

	return set[pos:]
}

type errorString string

func (e errorString) Error() string {
	return string(e)
}