package slicr

func Prepend[T any](values []T, val ...T) []T {
	values = append(val, values...)
	return values
}
