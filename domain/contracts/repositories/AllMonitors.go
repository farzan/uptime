package repositories

import "UptimeMonitor/domain/types/monitor"

type AllMonitors interface {
	Get() []monitor.Monitor
}
