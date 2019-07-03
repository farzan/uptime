package domain

import (
	"time"
)

type ResponseEntity struct {
	Id int
	MonitorId int
	Start time.Time
	End time.Time
	StatusCode int
	Error string
}

