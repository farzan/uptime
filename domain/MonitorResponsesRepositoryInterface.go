package domain

type MonitorResponsesRepositoryInterface interface {
	Store(response ResponseEntity)
	Find(id int) (ResponseEntity, error)
	FindByFilter(filter ResponseFilter) []ResponseEntity
}