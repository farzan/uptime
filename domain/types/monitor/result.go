package monitor

type Result struct {
	Monitor    Monitor
	ResultCode int
	Timeout    bool
	Duration   float64
}
