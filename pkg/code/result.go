package code

// Result is a generic type that represents the result of an request.
type Result[T any] struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// Success returns a result with code 200 and message "success"
func Success[T any](data T) Result[T] {
	return Result[T]{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

// Error returns a result with code 500 and the given message
func Error[T any](msg string, data T) Result[T] {
	return Result[T]{
		Code: 500,
		Msg:  msg,
		Data: data,
	}
}

// InError returns a result with code 500 and message "internal error"
func InError() Result[string] {
	return Result[string]{
		Code: 500,
		Msg:  "internal error",
		Data: "error",
	}
}
