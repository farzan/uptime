package response

import (
	"UptimeMonitor/domain"
	debug "UptimeMonitor/utils"
)

type Terminal struct {

}

func (t Terminal) Notify(response domain.Response) {
	debug.Printf("[%v] Response received; Dur: %v\n",
		response.Request.Monitor.Id,
		response.End.Sub(response.Start),
	)
}
