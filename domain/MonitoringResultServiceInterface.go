package domain

type MonitoringResultServiceInterface interface {
	StoreMonitoringResponse(response Response)
	StoreResponseEntity(response ResponseEntity)
	FindResponseById(id int) (Response, error)
	FindResponsesByFilter(filter ResponseFilter) []Response
	DeleteAll(monitorId int)
}
