package slicr

func Count[T any](values []T, predicate func(T) bool) int {
	var count int

	for _, val := range values {
		if predicate(val) {
			count += 1
		}
	}

	return count
}
