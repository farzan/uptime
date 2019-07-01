package monitor

type Thresholds interface {
	GetOk() int
	GetWarning() int
	GetCritical() int
}

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

func (t thresholds) GetOk() int {
	return t.Ok
}

func (t thresholds) GetWarning() int {
	return t.Warning
}

func (t thresholds) GetCritical() int {
	return t.Critical
}


