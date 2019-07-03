package domain

type Monitor struct {
	Id       int
	UserId   int
	Title    string
	Method   Method
	Url      string
	Interval int

	Thresholds
}

//func (m Monitor) IsResponseOk(r entities.Response) bool {
//	return r.Duration() < m.Thresholds.GetWarning()
//}
//
//func (m Monitor) IsResponseWarning(r entities.Response) bool {
//	return r.Duration() >= m.Thresholds.GetWarning() &&
//		r.Duration() < m.Thresholds.GetCritical()
//}
//
//func (m Monitor) IsResponseCritical(r entities.Response) bool {
//	return r.Duration() >= m.Thresholds.GetCritical()
//}
