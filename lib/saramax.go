package lib

type Handler[T any] struct {
	logger Logger
	fn     func() error
}
