package serviceProviders

import (
	repositories2 "UptimeMonitor/dataAccess/repositories"
	"UptimeMonitor/domain"
)

var (
	monitoringResultService domain.MonitoringResultServiceInterface
	allMonitorsRepository   repositories2.AllMonitorsRepositoryInterface
	userMonitorsRepository  repositories2.UserMonitorsRepositoryInterface
)

func init() {
	monitoringResultService = domain.NewMonitoringResultsServiceFake()
	allMonitorsRepository = domain.NewAllMonitorsRepositoryFake()
	userMonitorsRepository = repositories2.NewUserMonitorsRepositoryFake(0) // todo refactor
}

func GetMonitoringResultService() domain.MonitoringResultServiceInterface {
	return monitoringResultService
}

func GetAllMonitorsRepository() repositories2.AllMonitorsRepositoryInterface {
	return allMonitorsRepository
}

func GetUserMonitorsRepository() repositories2.UserMonitorsRepositoryInterface {
	return userMonitorsRepository
}
