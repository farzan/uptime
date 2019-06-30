package response

import (
	"UptimeMonitor/domain/types/monitor"
	debug "UptimeMonitor/utils"
)

type Terminal struct {

}

func (t Terminal) Notify(response monitor.Response) {
	debug.Printf("[%v] Response received; Dur: %v\n",
		response.Request.Monitor.Id,
		response.End.Sub(response.Start),
	)
}
