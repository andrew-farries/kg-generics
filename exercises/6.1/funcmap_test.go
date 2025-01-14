package funcmap_test

import (
	"fmt"
	"funcmap"
	"testing"
	"unicode"

	"github.com/google/go-cmp/cmp"
)

func TestFuncMap_AppliesDoubleTo2AndReturns4(t *testing.T) {
	t.Parallel()
	fm := funcmap.FuncMap[int, int]{
		"double": func(i int) int {
			return i * 2
		},
		"addOne": func(i int) int {
			return i + 1
		},
	}
	want := 4
	got := fm.Apply("double", 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFuncMap_AppliesUpperToUppercaseInputAndReturnsTrue(t *testing.T) {
	t.Parallel()
	fm := funcmap.FuncMap[rune, bool]{
		"upper": unicode.IsUpper,
		"lower": unicode.IsLower,
	}
	got := fm.Apply("upper", 'A')
	if !got {
		t.Fatal("upper('A'): want true, got false")
	}
}

func TestReverse(t *testing.T) {
	xs := []int{1, 2, 3}

	funcmap.Reverse(xs)

	fmt.Println(xs)
}

func TestMap(t *testing.T) {
	xs := []int{1, 2, 3}

	ys := funcmap.Map(xs, func(x int) int { return x + 1 })

	fmt.Println(ys)
}

func TestFilter(t *testing.T) {
	xs := []int{1, 2, 3}
	ys := []int32{1, 2, 3}

	zs1 := funcmap.Filter(xs, funcmap.IsEven[int])
	zs2 := funcmap.Filter(ys, funcmap.IsEven[int32])

	fmt.Println(zs1)
	fmt.Println(zs2)
}

func TestReduce(t *testing.T) {
	xs := []int{1, 2, 3}

	ys := funcmap.Reduce(xs, func(acc, curr int) int {
		return acc + curr
	}, 0)

	fmt.Println(ys)
}

func TestConcurrentFilter(t *testing.T) {
	xs := []int{1, 2, 3}

	ch := funcmap.ConcurrentFilter(xs, funcmap.IsEven[int])

	for x := range ch {
		fmt.Println(x)
	}
}
