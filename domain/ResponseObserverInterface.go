package domain

type ResponseObserver interface {
	Notify(response Response)
}