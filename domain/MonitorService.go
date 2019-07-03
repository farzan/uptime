package domain

import (
	repositories2 "UptimeMonitor/dataAccess/repositories"
)

type monitorService struct {

}

var monitorServiceInstance MonitorServiceInterface
var allMonitorsRepository repositories2.AllMonitorsRepositoryInterface

func init() {
	monitorServiceInstance = monitorService{}
	allMonitorsRepository = NewAllMonitorsRepositoryFake()
}

func GetMonitorService() MonitorServiceInterface {
	return monitorServiceInstance
}

func (m monitorService) Get(monitorId int) Monitor {
	return allMonitorsRepository.Get(monitorId)
}

func (m monitorService) GetAll() []Monitor {
	return allMonitorsRepository.All()
}

func (m monitorService) GetUserMonitors(userId int) []Monitor {
	repo := repositories2.NewUserMonitorsRepositoryFake(userId)
	return repo.All()
}
