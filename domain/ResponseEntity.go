package domain

type ResponseEntity struct {
	Id int
	MonitorId int
	Start string
	End string
	StatusCode int
	Error string
	ThresholdOk int
	ThresholdWarning int
	ThresholdCritical int
}

