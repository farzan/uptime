package entities

import (
	"UptimeMonitor/domain/types/monitor"
	"UptimeMonitor/serviceProviders"
	"time"
)

type Response struct {
	Id int
	MonitorId int
	Start time.Time
	End time.Time
	StatusCode int
	Error string

	monitor monitor.Monitor
}

func (r Response) Duration() time.Duration {
	return r.End.Sub(r.Start)
}

func (r Response) IsOk() bool {
	return r.Duration() < r.getMonitor().Thresholds.GetWarning()
}

func (r Response) IsWarning() bool {
	return r.Duration() >= r.getMonitor().Thresholds.GetWarning() &&
		r.Duration() < r.getMonitor().Thresholds.GetCritical()
}

func (r Response) IsCritical() bool {
	return r.Duration() >= r.getMonitor().Thresholds.GetCritical()
}

func (r Response) IsError() bool {
	return r.StatusCode < 200 || r.StatusCode > 299
}

func (r Response) getMonitor() monitor.Monitor {
	//todo Refactor
	return serviceProviders.GetUserMonitorsRepository().Get(r.MonitorId)
}
