package repositories

import (
	"UptimeMonitor/domain/types/monitor"
)

type UserMonitorsFake struct {
	monitors []monitor.Monitor
}

func (m UserMonitorsFake) Get() []monitor.Monitor {
	m.monitors = []monitor.Monitor{
		{1, 1, "My site 1", monitor.GET, "http://mysite1.com", 5, monitor.NewDefaultThreshold()},
		{2, 1, "My site 2", monitor.GET, "http://mysite2.com", 10, monitor.NewDefaultThreshold()},
		{3, 1, "My site 3", monitor.GET, "http://mysite3.com", 30, monitor.NewDefaultThreshold()},
		{4, 1, "My site 4", monitor.GET, "http://mysite4.com", 60, monitor.NewDefaultThreshold()},
	}

	return m.monitors
}

func (m UserMonitorsFake) Count() int {
	return len(m.monitors)
}

func (m UserMonitorsFake) Add(monitor monitor.Monitor) {
	m.monitors = append(m.monitors, monitor)
}

func (m UserMonitorsFake) Update(monitor monitor.Monitor) {
	for i, mon := range m.monitors {
		if monitor.Id == mon.Id {
			m.monitors[i] = monitor

			return
		}
	}

	// todo Raise error
}

func (m UserMonitorsFake) Delete(idOrModel interface{}) {
	if id, ok := idOrModel.(int); ok {
		m.deleteById(id)
	} else if model, ok := idOrModel.(monitor.Monitor); ok {
		m.deleteById(model.Id)
	}

	// todo Raise error
}

func (m UserMonitorsFake) deleteById(id int) {
	for index, monitor := range m.monitors {
		if monitor.Id == id {
			m.monitors = append(m.monitors[:index], m.monitors[index + 1:]...)

			return
		}
	}
}
