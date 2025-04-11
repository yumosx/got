package code

type Result[T any] struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func SuccessResponse[T any](data T) *Result[T] {
	return &Result[T]{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

func ErrorResponse[T any](msg string, data T) *Result[T] {
	return &Result[T]{
		Code: 500,
		Msg:  msg,
		Data: data,
	}
}
