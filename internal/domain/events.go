package domain

type Event struct {
	Key       interface{}
	Message   interface{}
	Operation string
}
