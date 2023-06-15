package merge_test

import (
	"merge"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergeCorrectlyMergesTwoMapsOfIntToBool(t *testing.T) {
	t.Parallel()
	inputs := []map[int]bool{
		{
			1: false,
			2: false,
			3: false,
		},
		{
			3: true,
			5: true,
		},
	}
	want := map[int]bool{
		1: false,
		2: false,
		3: true,
		5: true,
	}
	got := merge.Merge(inputs...)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestMergeCorrectlyMergesThreeMapsOfStringToAny(t *testing.T) {
	t.Parallel()
	inputs := []map[string]any{
		{
			"a": nil,
		},
		{
			"b": "hello, world",
			"c": 0,
		},
		{
			"a": 6 + 2i,
		},
	}
	want := map[string]any{
		"a": 6 + 2i,
		"b": "hello, world",
		"c": 0,
	}
	got := merge.Merge(inputs...)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestMergeMyMap(t *testing.T) {
	t.Parallel()
	m1 := map[string]int{"a": 2}
	m2 := merge.MyMap{"b": 3}

	want := map[string]int{
		"a": 2,
		"b": 3,
	}
	got := merge.Merge(m1, m2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDoSomething(t *testing.T) {
	t.Parallel()

	merge.DoSomething([]int{1, 2, 3})
	merge.DoSomething([]string{"foo", "bar", "baz"})

	xs := merge.MySlice{10, 11, 12}
	merge.DoSomething(xs)
}
