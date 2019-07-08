package domain

type UserMonitorsRepositoryFake struct {
	monitors []Monitor
}

func NewUserMonitorsRepositoryFake(userId int) *UserMonitorsRepositoryFake {
	monitors := getUserMonitors(userId)

	return &UserMonitorsRepositoryFake{monitors: monitors}
}

func getUserMonitors(userId int) []Monitor {
	monitors := GetAllMonitorsRepository().All()
	userMons := []Monitor{}
	for _, mon := range monitors {
		if mon.UserId == userId {
			userMons = append(userMons, mon)
		}
	}

	return userMons
}

func (m UserMonitorsRepositoryFake) All() []Monitor {
	return m.monitors
}

func (m UserMonitorsRepositoryFake) Get(id int) Monitor {
	return m.monitors[id]
}

func (m UserMonitorsRepositoryFake) Count() int {
	return len(m.monitors)
}

func (m UserMonitorsRepositoryFake) Add(monitor Monitor) {
	m.monitors = append(m.monitors, monitor)
}

func (m UserMonitorsRepositoryFake) Update(monitor Monitor) {
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
	} else if model, ok := idOrModel.(Monitor); ok {
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

func (m UserMonitorsRepositoryFake) UrlIsUnique(url string) bool {
	for _, mon := range m.monitors {
		if mon.Url == url {
			return false
		}
	}

	return true
}
