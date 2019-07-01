package entities

import "time"

type Response struct {
	Id int
	MonitorId int
	Start time.Time
	End time.Time
	StatusCode int
	Error string
}
