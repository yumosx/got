package stream

func Filter[T any](array []T, fn func(value T) bool) []T {
	var newArray = make([]T, len(array))
	for index, value := range newArray {
		if fn(value) {
			newArray[index] = array[index]
		}
	}
	return newArray
}
