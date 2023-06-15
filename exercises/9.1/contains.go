package contains

import "golang.org/x/exp/slices"

func ContainsFunc[T any](xs []T, fn func(T) bool) bool {
	return slices.IndexFunc(xs, fn) != -1
}
