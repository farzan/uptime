package domain

import (
	"time"
)

func NewMonitoringResultsService() MonitoringResultServiceInterface {
	s := MonitoringResultService{}
	s.responses = make(map[int] []ResponseEntity)

	return s
}

type MonitoringResultService struct {
	responses map[int] []ResponseEntity
}

func (s MonitoringResultService) StoreMonitoringResponse(response Response) {
	responseEntity := ResponseEntity{
		MonitorId: response.Request.Monitor.Id,
		Start: response.Start.Format(time.RFC3339Nano),
		End: response.End.Format(time.RFC3339Nano),
		StatusCode: response.StatusCode,
		Error: response.Error,
	}

	s.StoreResponseEntity(responseEntity)
}

func (s MonitoringResultService) StoreResponseEntity(response ResponseEntity) {
	GetMonitoringResponsesRepository().Store(response)
}

func (s MonitoringResultService) FindResponseById(id int) (Response, error) {
	responseEntity, _ := GetMonitoringResponsesRepository().Find(id)
	return s.createResponse(responseEntity), errorString("")
}

func (s MonitoringResultService) FindResponsesByFilter(filter ResponseFilter) []Response {
	responseEntities := GetMonitoringResponsesRepository().FindByFilter(filter)

	var responses []Response
	for _, responseEntity := range responseEntities {
		responses = append(responses, s.createResponse(responseEntity))
	}

	return responses
}

func (s MonitoringResultService) createResponse(responseEntity ResponseEntity) Response {
	start, _ := time.Parse(time.RFC3339Nano, responseEntity.Start)
	end, _ := time.Parse(time.RFC3339Nano, responseEntity.End)

	return Response{
		MonitorId: responseEntity.MonitorId,
		Start: start,
		End: end,
		StatusCode: responseEntity.StatusCode,
		Error: responseEntity.Error,
	}
}

type errorString string

func (e errorString) Error() string {
	return string(e)
}