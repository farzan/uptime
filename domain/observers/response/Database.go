package response

import (
	serviceContracts "UptimeMonitor/domain/contracts/services"
	"UptimeMonitor/domain/types/monitor"
	"UptimeMonitor/serviceProviders"
)

var service serviceContracts.MonitoringResultService

func init() {
	service = serviceProviders.GetMonitoringResultService()
}

type Database struct {

}

func (d Database) Notify(response monitor.Response) {
	service.StoreMonitoringResponse(response)
}
