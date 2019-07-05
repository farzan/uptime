package domain

type MonitorResponsesRepositoryInterface interface {
	Store(response Response)
	Find(id int) Response
	FindByFilter(filter ResponseFilter) []Response
}