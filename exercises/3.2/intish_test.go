package intish_test

import (
	"fmt"
	"intish"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type MyInt int

func TestIsPositive_IsTrueFor1(t *testing.T) {
	t.Parallel()
	input := MyInt(1)
	want := true
	got := intish.IsPositive(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsPositive_IsFalseForNegative1(t *testing.T) {
	t.Parallel()
	input := MyInt(-1)
	want := false
	got := intish.IsPositive(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsPositive_IsFalseForZero(t *testing.T) {
	t.Parallel()
	input := MyInt(0)
	want := false
	got := intish.IsPositive(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFoo(t *testing.T) {
	x := intish.MyCoolInt(12)

	fmt.Println(intish.Foo(x))
}

func TestContains(t *testing.T) {
	xs := []int{1, 2, 3}

	actual := intish.Contains(xs, 2)

	if !cmp.Equal(actual, true) {
		t.Error(cmp.Diff(true, actual))
	}
}

func TestContainsFn(t *testing.T) {
	xs := []intish.AnotherInt{
		intish.AnotherInt(1),
		intish.AnotherInt(2),
		intish.AnotherInt(3),
	}

	actual := intish.ContainsFn(xs, 2)

	if !cmp.Equal(actual, true) {
		t.Error(cmp.Diff(true, actual))
	}
}

func TestContainsFn2(t *testing.T) {
	xs := []intish.AnotherInt{
		intish.AnotherInt(1),
		intish.AnotherInt(2),
		intish.AnotherInt(3),
	}

	actual := intish.ContainsFn2(xs, 2)

	if !cmp.Equal(actual, true) {
		t.Error(cmp.Diff(true, actual))
	}
}
