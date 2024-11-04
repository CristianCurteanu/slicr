package slicr

func Map[I any, O any](input []I, predicate func(I) O) []O {
	var output []O = make([]O, 0, len(input))
	for _, i := range input {
		output = append(output, predicate(i))
	}

	return output
}
