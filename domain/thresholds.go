package domain

import "time"

type Thresholds interface {
	GetOk() time.Duration
	GetWarning() time.Duration
	GetCritical() time.Duration
}

type thresholds struct {
	Ok time.Duration
	Warning time.Duration
	Critical time.Duration
}

func NewThreshold(ok int, warning int, critical int) Thresholds {
	return &thresholds{
		time.Duration(ok) * time.Second,
		time.Duration(warning) * time.Second,
		time.Duration(critical) * time.Second,
	}
}

func NewDefaultThreshold() Thresholds {
	return thresholds{
		time.Duration(1) * time.Second,
		time.Duration(3) * time.Second,
		time.Duration(7) * time.Second,
	}
}

func (t thresholds) GetOk() time.Duration {
	return t.Ok
}

func (t thresholds) GetWarning() time.Duration {
	return t.Warning
}

func (t thresholds) GetCritical() time.Duration {
	return t.Critical
}


