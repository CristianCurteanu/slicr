package slicr

func Contains[T comparable](values []T, target T) bool {
	_, found := Find(values, func(tv T) bool { return tv == target })
	return found
}
