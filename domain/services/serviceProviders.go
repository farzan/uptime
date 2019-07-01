package services

import (
	serviceContracts "UptimeMonitor/domain/contracts/services"
)

var monitoringResultService serviceContracts.MonitoringResultService

func init() {
	monitoringResultService = NewMonitoringResultsServiceFake()
}

func GetMonitoringResultService() serviceContracts.MonitoringResultService {
	return monitoringResultService
}
