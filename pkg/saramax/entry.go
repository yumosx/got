package saramax

// Message 表示发送到消息体
type Message[T any] struct {
	data T
}

func (e Message[T]) Encode() ([]byte, error) {
	return nil, nil
}

func (e Message[T]) Length() int {
	return 0
}

func NewEntry[T any](data T) Message[T] {
	return Message[T]{data: data}
}
