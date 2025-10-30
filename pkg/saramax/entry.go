package saramax

import (
	"encoding/json"
)

type Message[T any] struct {
	data T
}

func (e Message[T]) Encode() ([]byte, error) {
	return json.Marshal(e.data)
}

func (e Message[T]) Length() int {
	data, err := e.Encode()
	if err != nil {
		return 0
	}
	return len(data)
}

func (e Message[T]) Data() T {
	return e.data
}

func NewEntry[T any](data T) Message[T] {
	return Message[T]{data: data}
}
