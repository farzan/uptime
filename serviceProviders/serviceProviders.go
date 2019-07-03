package serviceProviders

import (
	repositories2 "UptimeMonitor/dataAccess/repositories"
	"UptimeMonitor/domain/contracts/repositories"
	serviceContracts "UptimeMonitor/domain/contracts/services"
	"UptimeMonitor/domain/services"
)

var (
	monitoringResultService serviceContracts.MonitoringResultService
	allMonitorsRepository repositories.AllMonitorsRepository
	userMonitorsRepository repositories.UserMonitorsRepository
)

func init() {
	monitoringResultService = services.NewMonitoringResultsServiceFake()
	allMonitorsRepository = repositories2.NewAllMonitorsRepositoryFake()
	userMonitorsRepository = repositories2.NewUserMonitorsRepositoryFake()
}

func GetMonitoringResultService() serviceContracts.MonitoringResultService {
	return monitoringResultService
}

func GetAllMonitorsRepository() repositories.AllMonitorsRepository {
	return allMonitorsRepository
}

func GetUserMonitorsRepository() repositories.UserMonitorsRepository {
	return userMonitorsRepository
}
