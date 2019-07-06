package domain

import debug "UptimeMonitor/utils"

type AllMonitorsRepositoryFake struct {
	monitors []Monitor
}

func (m AllMonitorsRepositoryFake) All() []Monitor {
	return m.monitors
}

func (m AllMonitorsRepositoryFake) Get(monitorId int) (Monitor, error) {
	for _, mon := range m.monitors {
		if mon.Id == monitorId {
			return mon, nil
		}
	}

	return Monitor{}, errorString("Monitor not found")
}

func NewAllMonitorsRepositoryFake() *AllMonitorsRepositoryFake {
	monitors := []Monitor{
		{1, 1, "My site 1", GET, "http://mysite1.com", 5, NewDefaultThreshold()},
		{2, 1, "My site 2", GET, "http://mysite2.com", 10, NewDefaultThreshold()},
		{3, 1, "My site 3", GET, "http://mysite3.com", 30, NewDefaultThreshold()},
		{4, 1, "My site 4", GET, "http://mysite4.com", 60, NewDefaultThreshold()},
	}

	return &AllMonitorsRepositoryFake{monitors: monitors}
}

func (m *AllMonitorsRepositoryFake) Add(monitor *Monitor) {
	monitor.Id = m.getNewId()
	m.monitors = append(m.monitors, *monitor)
	debug.Printf("len:%v\n", len(m.monitors))
}

func (m *AllMonitorsRepositoryFake) getNewId() int {
	id := 0
	for _, mon := range m.monitors {
		if mon.Id > id {
			id = mon.Id
		}
	}

	return id + 1
}

func(m AllMonitorsRepositoryFake) Update(monitor *Monitor) error {
	// todo
	return errorString("Not found")
}

func (m AllMonitorsRepositoryFake) Delete(monitorId int) error {
	for i, mon := range m.monitors {
		if mon.Id == monitorId {
			m.monitors = append(m.monitors[:i], m.monitors[i+1:]...)
			return nil
		}
	}
	return errorString("Not found")
}
