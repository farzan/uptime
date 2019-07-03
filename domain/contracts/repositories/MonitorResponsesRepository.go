package repositories

import (
	"UptimeMonitor/domain/entities"
	"UptimeMonitor/domain/services/filters"
)

type MonitorResponsesRepository interface {
	Store(response entities.Response)
	Find(id int) entities.Response
	FindByFilter(filter filters.ResponseFilter) []entities.Response
}