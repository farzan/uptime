package repositories

import (
	"UptimeMonitor/domain"
)

type UserMonitorsRepositoryInterface interface {
	All() []domain.Monitor
	Get(id int) domain.Monitor
	Count() int
	Add(monitor domain.Monitor)
	Update(monitor domain.Monitor)
	Delete(idOrModel interface{})
}
