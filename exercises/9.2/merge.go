package merge

import "golang.org/x/exp/maps"

func Merge[K comparable, T any](ms ...map[K]T) map[K]T {
	dst := map[K]T{}

	for _, m := range ms {
		maps.Copy(dst, m)
	}
	return dst
}
