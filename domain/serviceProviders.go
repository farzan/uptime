package domain

var (
	monitoringResultService MonitoringResultServiceInterface
	allMonitorsRepository   AllMonitorsRepositoryInterface
	userMonitorsRepository  UserMonitorsRepositoryInterface
	responsesRepository     MonitorResponsesRepositoryInterface
	workerServiceInstance	WorkerServiceInterface
)

func init() {
	userMonitorsRepository = NewUserMonitorsRepositoryFake(0) // todo refactor
}

func GetMonitoringResultService() MonitoringResultServiceInterface {
	if monitoringResultService == nil {
		monitoringResultService = NewMonitoringResultsService()
	}
	return monitoringResultService
}

func GetAllMonitorsRepository() AllMonitorsRepositoryInterface {
	if allMonitorsRepository == nil {
		allMonitorsRepository = NewAllMonitorsRepositoryFake()
	}
	return allMonitorsRepository
}

func GetUserMonitorsRepository(userId int) UserMonitorsRepositoryInterface {
	return userMonitorsRepository
}

func GetMonitoringResponsesRepository() MonitorResponsesRepositoryInterface {
	if responsesRepository == nil {
		responsesRepository = NewMonitoringResponsesRepositoryFake()
	}
	return responsesRepository;
}

func GetWorkerService() WorkerServiceInterface {
	if workerServiceInstance == nil {
		workerServiceInstance = NewWorkerService()
	}
	return  workerServiceInstance
}
