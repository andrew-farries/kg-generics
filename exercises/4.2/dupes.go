package dupes

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func Dupes[T comparable](xs []T) bool {
	m := make(map[T]bool, len(xs))

	for _, x := range xs {
		if _, ok := m[x]; ok {
			return true
		}
		m[x] = true
	}
	return false
}

func Max[T constraints.Ordered](xs []T) (T, error) {
	var max T

	if len(xs) == 0 {
		return max, errors.New("cannot find the maximum value of an empty slice")
	}

	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max, nil
}
