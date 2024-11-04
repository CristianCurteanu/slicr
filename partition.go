package slicr

func Partition[T any](values []T, predicate func(T) bool) ([]T, []T) {
	var trues []T = make([]T, 0, len(values))
	var falses []T = make([]T, 0, len(values))

	for _, val := range values {
		if predicate(val) {
			trues = append(trues, val)
		} else {
			falses = append(falses, val)
		}
	}

	return trues, falses
}
