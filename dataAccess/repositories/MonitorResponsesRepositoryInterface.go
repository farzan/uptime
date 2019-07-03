package repositories

import (
	"UptimeMonitor/domain"
)

type MonitorResponsesRepositoryInterface interface {
	Store(response domain.Response)
	Find(id int) domain.Response
	FindByFilter(filter domain.ResponseFilter) []domain.Response
}