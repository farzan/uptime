package domain

const MaxUserMonitors = 5

type monitorService struct {

}

func GetMonitorService() MonitorServiceInterface {
	return monitorService{}
}

func (m monitorService) Find(monitorId int) (Monitor, error) {
	return GetAllMonitorsRepository().Get(monitorId)
}

func (m monitorService) FindAll() []Monitor {
	return GetAllMonitorsRepository().All()
}

func (m monitorService) FindUserMonitors(userId int) []Monitor {
	return GetUserMonitorsRepository(userId).All()
}

func (m monitorService) Add(monitor *Monitor) {
	GetAllMonitorsRepository().Add(monitor)
}

func (m monitorService) Update(monitor *Monitor) error {
	return GetAllMonitorsRepository().Update(monitor)
}

func (m monitorService) Delete(monitorId int) error {
	return GetAllMonitorsRepository().Delete(monitorId)
}

func (m monitorService) UrlIsUnique(userId int, url string) bool {
	return GetUserMonitorsRepository(userId).UrlIsUnique(url)
}

func (m monitorService) UserHasReachedMax(userId int) bool {
	return GetUserMonitorsRepository(userId).Count() <= MaxUserMonitors
}
