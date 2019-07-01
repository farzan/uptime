package services

import (
	"UptimeMonitor/domain/entities"
	"UptimeMonitor/domain/services/filters"
	"UptimeMonitor/domain/types/monitor"
)

type MonitoringResultService interface {
	StoreMonitoringResponse(response monitor.Response)
	StoreResponse(response entities.Response)
	FindResponseById(id int) (entities.Response, error)
	FindResponsesByFilter(filter filters.ResponseFilter) []entities.Response
}
