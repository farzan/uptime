package domain

type UserMonitorsRepositoryInterface interface {
	All() []Monitor
	Get(id int) Monitor
	Count() int
	Add(monitor Monitor)
	Update(monitor Monitor)
	Delete(idOrModel interface{})
	UrlIsUnique(url string) bool
}
