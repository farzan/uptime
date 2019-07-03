package repositories

import "UptimeMonitor/domain/entities"

type ResultsRepositoryFake struct {

}

func (r ResultsRepositoryFake) FindAll(monitorId int, offset int, count int) []entities.Response {

}

func (r ResultsRepositoryFake) Create(entities.Response) {

}
