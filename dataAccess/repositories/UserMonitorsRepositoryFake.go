package repositories

import (
	"UptimeMonitor/domain"
)

type UserMonitorsRepositoryFake struct {
	monitors []domain.Monitor
}

func NewUserMonitorsRepositoryFake(userId int) *UserMonitorsRepositoryFake {
	monitors := domain.NewAllMonitorsRepositoryFake().monitors // Dont extract

	return &UserMonitorsRepositoryFake{monitors: monitors}
}

func (m UserMonitorsRepositoryFake) All() []domain.Monitor {
	return m.monitors
}

func (m UserMonitorsRepositoryFake) Get(id int) domain.Monitor {
	return m.monitors[id]
}

func (m UserMonitorsRepositoryFake) Count() int {
	return len(m.monitors)
}

func (m UserMonitorsRepositoryFake) Add(monitor domain.Monitor) {
	m.monitors = append(m.monitors, monitor)
}

func (m UserMonitorsRepositoryFake) Update(monitor domain.Monitor) {
	for i, mon := range m.monitors {
		if monitor.Id == mon.Id {
			m.monitors[i] = monitor

			return
		}
	}

	// todo Raise error
}

func (m UserMonitorsRepositoryFake) Delete(idOrModel interface{}) {
	if id, ok := idOrModel.(int); ok {
		m.deleteById(id)
	} else if model, ok := idOrModel.(domain.Monitor); ok {
		m.deleteById(model.Id)
	}

	// todo Raise error
}

func (m UserMonitorsRepositoryFake) deleteById(id int) {
	for index, mon := range m.monitors {
		if mon.Id == id {
			m.monitors = append(m.monitors[:index], m.monitors[index + 1:]...)

			return
		}
	}
}
