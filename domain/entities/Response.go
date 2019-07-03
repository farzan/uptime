package entities

import (
	"UptimeMonitor/dataAccess/repositories"
	"UptimeMonitor/domain/types/monitor"
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

func (r Response) Duration() time.Duration {
	return r.End.Sub(r.Start)
}

func (r Response) getMonitor() monitor.Monitor {
	//todo Refactor
	return repositories.NewUserMonitorsFake().Get(r.MonitorId)
}