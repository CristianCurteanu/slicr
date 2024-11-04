package slicr

func Filter[T any](values []T, predicate func(T) bool) []T {
	/*
		This trick uses the fact that a slice shares the same backing array and capacity as the original, so the storage is reused for the filtered slice.
		This way it won't allocate new slice
	*/
	var selected []T = values[:0]
	for _, el := range values {
		if predicate(el) {
			selected = append(selected, el)
		}
	}

	// For elements which must be garbage collected, it will set it's zero value
	var zv T
	for i := len(selected); i < len(values); i++ {
		values[i] = zv // or the zero value of T
	}

	return selected
}
