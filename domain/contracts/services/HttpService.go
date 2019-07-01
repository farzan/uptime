package domain

import "UptimeMonitor/domain/types/monitor"

type HttpService interface {
	Process(request monitor.Request) monitor.Response
}
