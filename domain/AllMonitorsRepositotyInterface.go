package domain

type AllMonitorsRepositoryInterface interface {
	All() []Monitor
	Get(monitorId int) (Monitor, error)
	Add(monitor *Monitor)
	Update(monitor *Monitor) error
	Delete(monitorId int) error
}
