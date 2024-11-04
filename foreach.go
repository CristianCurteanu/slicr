package slicr

func ForEach[I any](values []I, predicate func(I)) {
	for _, i := range values {
		predicate(i)
	}
}
