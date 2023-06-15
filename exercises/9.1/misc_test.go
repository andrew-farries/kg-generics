package contains_test

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSlices(t *testing.T) {
	xs := []string{"a", "b"}
	ys := []string{"A", "B"}

	fmt.Println(slices.EqualFunc(xs, ys, strings.EqualFold))

	zs := []int{3, 2, 1}
	slices.Sort(zs)
	fmt.Println(zs)
	fmt.Println(slices.Clone(zs))
}
