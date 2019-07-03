package domain

type MonitoringResultServiceInterface interface {
	StoreMonitoringResponse(response Response)
	StoreResponse(response Response)
	FindResponseById(id int) (Response, error)
	FindResponsesByFilter(filter ResponseFilter) []Response
}
