package repositories

import "UptimeMonitor/domain/types/monitor"

type AllMonitorsRepository interface {
	All() []monitor.Monitor
}
