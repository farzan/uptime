package domain

type ResultsRepositoryInterface interface {
	FindAll(monitorId int, offset int, count int) []Response
	Create(Response)
}
