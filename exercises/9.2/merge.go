package merge

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func Merge[K comparable, T any](ms ...map[K]T) map[K]T {
	dst := map[K]T{}

	for _, m := range ms {
		maps.Copy(dst, m)
	}
	return dst
}

type MyMap map[string]int

type MySlice []int

func DoSomething[T any](xs []T) {
	fmt.Printf("doing something with %v\n", xs)
}
