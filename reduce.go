package slicr

func Reduce[T any](values []T, reducer func(T, T) T) T {
	var result T
	for i, val := range values {
		if i == 0 {
			result = val
			continue
		}
		result = reducer(result, val)
	}

	return result
}
