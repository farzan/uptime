package services

import (
	"UptimeMonitor/domain/types/monitor"
)

type ResponseObserver interface {
	Notify(response monitor.Response)
}