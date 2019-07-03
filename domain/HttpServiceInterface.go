package domain

type HttpServiceInterface interface {
	Process(request Request) Response
}
