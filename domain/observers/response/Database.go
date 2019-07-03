package response

import (
	"UptimeMonitor/domain"
)

var service = domain.NewMonitoringResultsServiceFake()

type Database struct {

}

func (d Database) Notify(response domain.Response) {
	service.StoreMonitoringResponse(response)
}
