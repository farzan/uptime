package repositories

import "UptimeMonitor/domain/types/monitor"

type AllMonitorsFake struct {
	monitors []monitor.Monitor
}

func (m AllMonitorsFake) All() []monitor.Monitor {
	return m.monitors
}

func NewAllMonitorsFake() *AllMonitorsFake {
	monitors := []monitor.Monitor{
		{1, 1, "My site 1", monitor.GET, "http://mysite1.com", 5, monitor.NewDefaultThreshold()},
		{2, 1, "My site 2", monitor.GET, "http://mysite2.com", 10, monitor.NewDefaultThreshold()},
		{3, 1, "My site 3", monitor.GET, "http://mysite3.com", 30, monitor.NewDefaultThreshold()},
		{4, 1, "My site 4", monitor.GET, "http://mysite4.com", 60, monitor.NewDefaultThreshold()},
	}

	return &AllMonitorsFake{monitors: monitors}
}