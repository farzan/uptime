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
	return int(r.Duration() * time.Second) < r.getMonitor().Thresholds.GetOk()
}

func (r Response) IsWarning() bool {
	dur := int(r.Duration() * time.Second)
	return dur >= r.getMonitor().Thresholds.GetOk() &&
		dur < r.getMonitor().Thresholds.GetWarning()
}

func (r Response) IsCritical() bool {
	dur := int(r.Duration() * time.Second)
	return dur >= r.getMonitor().Thresholds.GetCritical()
}

func (r Response) Duration() time.Duration {
	return r.End.Sub(r.Start)
}

func (r Response) DurationInSeconds() float64 {
	return float64(r.End.Sub(r.Start) / time.Second)
}

func (r Response) getMonitor() monitor.Monitor {
	//todo Refactor
	return repositories.NewUserMonitorsFake().Get(r.MonitorId)
}