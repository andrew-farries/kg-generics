package print_test

import (
	"bytes"
	"fmt"
	"print"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrintAnythingTo_PrintsInputToSuppliedWriter(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	print.PrintAnythingTo[string](buf, "Hello, world")
	want := "Hello, world\n"
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestLen(t *testing.T) {
	t.Parallel()

	got := print.Len([]int{1, 2, 3})
	want := 3

	if got != want {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDrain(t *testing.T) {
	t.Parallel()
	ch := make(chan int)

	go func() {
		defer close(ch)

		vals := []int{1, 2, 3, 4}
		for _, val := range vals {
			ch <- val
		}
	}()

	print.Drain(ch)
}

func TestMerge(t *testing.T) {
	t.Parallel()

	ch1, ch2 := make(chan int), make(chan int)

	sendTo := func(ch chan int, xs []int) {
		defer close(ch)

		for _, x := range xs {
			ch <- x
		}
	}

	out := print.Merge(ch1, ch2)

	go sendTo(ch1, []int{1, 2, 3})
	go sendTo(ch2, []int{4, 5, 6})

	for val := range out {
		fmt.Println(val)
	}
}

func TestBunch(t *testing.T) {
	t.Parallel()

	b := print.Bunch[int]{1, 2, 3}
	b = append(b, 4)

	fmt.Println(b)
}
