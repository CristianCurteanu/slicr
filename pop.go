package slicr

func Pop[T any](values []T) (T, []T) {
	return values[len(values)-1], values[:len(values)-1]
}
