package domain

type AllMonitorsRepositoryInterface interface {
	All() []Monitor
	Get(monitorId int) Monitor
}
