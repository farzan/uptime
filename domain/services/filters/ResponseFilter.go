package filters

type ResponseFilter interface {
	GetMonitorId() int
	GetCount() int
}

func NewResponseFilter(monitorId int, count int) ResponseFilter {
	return responseFilter{
		monitorId: monitorId,
		count: count,
	}
}

type responseFilter struct {
	monitorId int
	count     int
}

func (f responseFilter) GetMonitorId() int {
	return f.monitorId
}

func (f responseFilter) GetCount() int {
	return f.count
}
