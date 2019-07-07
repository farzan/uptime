package domain

import debug "UptimeMonitor/utils"

func NewMonitoringResponsesRepositoryFake() MonitorResponsesRepositoryInterface {
	s := monitorResponseRepositoryFake{}
	s.responses = make(map[int] []ResponseEntity)

	return &s
}

type monitorResponseRepositoryFake struct {
	responses map[int] []ResponseEntity
}

func (r *monitorResponseRepositoryFake) Store(response ResponseEntity) {
	r.responses[response.MonitorId] = append([]ResponseEntity{response}, r.responses[response.MonitorId]...)

	// Debug code:
	lengths := make(map[int]int)
	for k, v := range r.responses {
		lengths[k] = len(v)
	}
	debug.Printf("===> response stored; count: %v\n", lengths)
}

func (r monitorResponseRepositoryFake) Find(id int) (ResponseEntity, error) {
	// todo
	return ResponseEntity{}, errorString("")

}

func (r monitorResponseRepositoryFake) FindByFilter(filter ResponseFilter) []ResponseEntity {
	monId := filter.GetMonitorId()
	set := r.responses[monId]
	pos := len(set) - filter.GetCount()

	if pos < 0 {
		pos = 0
	}

	return set[pos:]
}

func (r monitorResponseRepositoryFake) DeleteAll(monitorId int) {
	delete(r.responses, monitorId)
}
