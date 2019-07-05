package domain

var (
	monitoringResultService MonitoringResultServiceInterface
	allMonitorsRepository   AllMonitorsRepositoryInterface
	userMonitorsRepository  UserMonitorsRepositoryInterface
)

func init() {
	userMonitorsRepository = NewUserMonitorsRepositoryFake(0) // todo refactor
}

func GetMonitoringResultService() MonitoringResultServiceInterface {
	if monitoringResultService == nil {
		monitoringResultService = NewMonitoringResultsServiceFake()
	}
	return monitoringResultService
}

func GetAllMonitorsRepository() AllMonitorsRepositoryInterface {
	if allMonitorsRepository == nil {
		allMonitorsRepository = NewAllMonitorsRepositoryFake()
	}
	return allMonitorsRepository
}

func GetUserMonitorsRepository() UserMonitorsRepositoryInterface {
	return userMonitorsRepository
}
