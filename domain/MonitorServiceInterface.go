package domain

type MonitorServiceInterface interface {
	Get(monitorId int) Monitor
	GetAll() []Monitor
	GetUserMonitors(userId int) []Monitor
	//AddMonitorForUser()
}
