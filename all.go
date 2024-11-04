package slicr

func All[T any](values []T, predicate func(T) bool) bool {
	for _, val := range values {
		if !predicate(val) {
			return false
		}
	}
	return true
}
