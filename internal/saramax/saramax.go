package saramax

import "github.com/yumosx/got/internal/synx"

type Handler[T any] struct {
	logger synx.Logger
	fn     func() error
}
