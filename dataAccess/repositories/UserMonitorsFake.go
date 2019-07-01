package repositories

import (
	"UptimeMonitor/domain/types/monitor"
)

type UserMonitorsFake struct {
	monitors []monitor.Monitor
}

func NewUserMonitorsFake() *UserMonitorsFake {
	monitors := NewAllMonitorsFake().monitors

	return &UserMonitorsFake{monitors: monitors}
}

func (m UserMonitorsFake) Get() []monitor.Monitor {
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
	for index, mon := range m.monitors {
		if mon.Id == id {
			m.monitors = append(m.monitors[:index], m.monitors[index + 1:]...)

			return
		}
	}
}
