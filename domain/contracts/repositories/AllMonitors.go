package repositories

import "UptimeMonitor/domain/types/monitor"

type AllMonitors interface {
	All() []monitor.Monitor
}
