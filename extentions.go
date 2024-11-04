package slicr

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Extend[T any](a []T, n int) []T {
	return append(a, make([]T, n)...)
}

func ExtendAt[T any](a []T, n, i int) []T {
	return append(a[:i], append(make([]T, n), a[i:]...)...)
}

func ExtendCap[T any](a []T, n int) []T {
	// Make sure there is space to append n elements without re-allocating
	if cap(a)-len(a) < n {
		a = append(make([]T, 0, len(a)+n), a...)
	}

	return a
}

func Deduplicate[T constraints.Ordered](values []T) []T {
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	j := 0
	for i := 1; i < len(values); i++ {
		if values[j] == values[i] {
			continue
		}
		j++
		values[j] = values[i]
	}
	result := values[:j+1]

	return result
}
