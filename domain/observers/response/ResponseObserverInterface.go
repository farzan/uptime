package response

import (
	"UptimeMonitor/domain"
)

type ResponseObserver interface {
	Notify(response domain.Response)
}