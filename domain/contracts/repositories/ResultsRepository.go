package repositories

import "UptimeMonitor/domain/entities"

type ResultsRepository interface {
	FindAll(monitorId int, offset int, count int) []entities.Response
	Create(entities.Response)
}
