package domain

type MonitorServiceInterface interface {
	Find(monitorId int) (Monitor, error)
	FindAll() []Monitor
	FindUserMonitors(userId int) []Monitor
	UrlIsUnique(userId int, url string) bool
	Add(monitor *Monitor)
	Update(monitor *Monitor) error
	Delete(monitorId int) error
	UserHasReachedMax(userId int) bool
}
