package monitor

type Monitor struct {
	Id         int
	UserId     int
	Title      string
	Method     Method
	Url        string
	Interval   int
	Thresholds Thresholds
}
