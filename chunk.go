package slicr

func Chunk[T any](values []T, size int) [][]T {
	i := len(values)
	chunks, remainder := i/size, i%size

	if remainder != 0 {
		chunks += 1
	}
	var res [][]T = make([][]T, chunks)
	for i := 0; i < chunks; i++ {
		res[i] = make([]T, 0, size)

		for j := 0; j < size; j++ {
			if len(values) == 0 {
				break
			}
			var temp T
			temp, values = values[0], values[1:]
			res[i] = append(res[i], temp)
		}
	}

	return res
}
