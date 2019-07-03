package domain

type AllMonitorsRepositoryFake struct {
	monitors []Monitor
}

func (m AllMonitorsRepositoryFake) All() []Monitor {
	return m.monitors
}

func (m AllMonitorsRepositoryFake) Get(monitorId int) Monitor {
	return m.monitors[monitorId]
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