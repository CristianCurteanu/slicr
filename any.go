package slicr

func Any[T any](values []T, predicate func(T) bool) bool {
	_, found := Find(values, predicate)
	return found
}
