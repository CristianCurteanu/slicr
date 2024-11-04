package slicr

func Find[T any](values []T, predicate func(T) bool) (T, bool) {
	for _, el := range values {
		if predicate(el) {
			return el, true
		}
	}
	var zeroVal T
	return zeroVal, false
}
