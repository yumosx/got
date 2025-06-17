package errx

import "fmt"

type Option[T any] struct {
	Val T
	err error
}

func Ok[T any](value T) Option[T] {
	return Option[T]{Val: value, err: nil}
}

// Err 默认使用0 值, return Err[string](xxx)
func Err[T any](err error) Option[T] {
	var zero T
	return Option[T]{Val: zero, err: err}
}

func Errf[T any](tmpl string, value string) Option[T] {
	var zero T
	return Option[T]{Val: zero, err: fmt.Errorf(tmpl, value)}
}

// VErr 我们希望用户去控制这个 value, return VErr(result{}, err)
func VErr[T any](val T, err error) Option[T] {
	return Option[T]{Val: val, err: err}
}

func (o Option[T]) Ret() (T, error) {
	return o.Val, o.err
}

func (o Option[T]) Error() error {
	return o.err
}

func (o Option[T]) NoNil() bool {
	return o.err != nil
}
