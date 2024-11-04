package slicr

func Delete[T any](a []T, i int) []T {
	var zeroVal T
	copy(a[i:], a[i+1:])
	a[len(a)-1] = zeroVal // the zero value of slice
	return a[:len(a)-1]
}
