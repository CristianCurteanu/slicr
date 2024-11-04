package slicr

func Shift[T any](values []T) (T, []T) {
	res, values := values[0], values[1:]
	return res, values
}
