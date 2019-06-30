package monitor

import "time"

type Response struct {
	Request Request
	Start time.Time
	End time.Time
	StatusCode int
	Error string
}

func (r Response) Duration() time.Duration {
	return r.End.Sub(r.Start)
}
