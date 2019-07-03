package repositories

import (
	"UptimeMonitor/domain"
)

type ResultsRepositoryInterface interface {
	FindAll(monitorId int, offset int, count int) []domain.Response
	Create(domain.Response)
}
