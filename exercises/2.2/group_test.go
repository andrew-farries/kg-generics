package group_test

import (
	"fmt"
	"group"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGroupContainsWhatIsAppendedToIt(t *testing.T) {
	t.Parallel()
	got := group.Group[string]{}
	got = append(got, "hello")
	got = append(got, "world")
	want := group.Group[string]{"hello", "world"}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIdentity(t *testing.T) {
	t.Parallel()

	// This works but you can't do partial application of types
	// (no currying)
	var fn group.IdFunc[int] = group.Identity[int]

	fmt.Printf("%T\n", fn)
}
