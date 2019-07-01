package response

import (
	serviceContracts "UptimeMonitor/domain/contracts/services"
	"UptimeMonitor/domain/services"
	"UptimeMonitor/domain/types/monitor"
)

var service serviceContracts.MonitoringResultService

func init() {
	service = services.GetMonitoringResultService()
}

type Database struct {

}

func (d Database) Notify(response monitor.Response) {
	service.StoreMonitoringResponse(response)
}
