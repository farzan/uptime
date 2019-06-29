package monitor

import "time"

type Response struct {
	Request Request
	Start time.Time
	End time.Time
	StatusCode int
	Error string
}
