package repositories

import (
	"UptimeMonitor/domain/types/monitor"
)

type UserMonitors interface {
	Get() []monitor.Monitor
	Count() int
	Add(monitor monitor.Monitor)
	Update(monitor monitor.Monitor)
	Delete(idOrModel interface{})
}
