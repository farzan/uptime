package domain

import (
	debug "UptimeMonitor/utils"
)

type TerminalObserver struct {

}

func (t TerminalObserver) Notify(response Response) {
	debug.Printf("[%v] Response received; Dur: %v\n",
		response.Request.Monitor.Id,
		response.End.Sub(response.Start),
	)
}
