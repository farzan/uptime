package monitor

type Thresholds interface {}

type thresholds struct {
	Ok int
	Warning int
	Critical int
}

func NewThreshold(ok int, warning int, critical int) Thresholds {
	return &thresholds{ok, warning, critical}
}

func NewDefaultThreshold() Thresholds {
	return thresholds{2000, 5000, 15000}
}


