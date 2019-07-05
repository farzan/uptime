package domain

var service MonitoringResultServiceInterface

func init() {
	service = GetMonitoringResultService()
}

type DatabaseObserver struct {

}

func (d DatabaseObserver) Notify(response Response) {
	service.StoreMonitoringResponse(response)
}
