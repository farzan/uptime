package repositories

import (
	"UptimeMonitor/domain"
)

type ResultsRepositoryFake struct {

}

func (r ResultsRepositoryFake) FindAll(monitorId int, offset int, count int) []domain.Response {

}

func (r ResultsRepositoryFake) Create(domain.Response) {

}
