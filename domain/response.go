package domain

import (
	"time"
)

type Response struct {
	MonitorId int
	Start time.Time
	End time.Time
	StatusCode int
	Error string

	Request Request
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

func (r Response) getMonitor() Monitor {
	monitor, _ := GetMonitorService().Find(r.MonitorId)
	return monitor
}