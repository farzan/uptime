package repositories

import (
	"UptimeMonitor/domain"
)

type AllMonitorsRepositoryInterface interface {
	All() []domain.Monitor
	Get(monitorId int) domain.Monitor
}
