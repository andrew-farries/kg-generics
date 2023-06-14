package sequence_test

import (
	"errors"
	"fmt"
	"sequence"
	"testing"
)

func TestEmptyIsTrueForEmptySequence(t *testing.T) {
	t.Parallel()
	s := sequence.Sequence[int]{}
	got := s.Empty()
	if !got {
		t.Fatal("false for empty sequence")
	}
}

func TestEmptyIsFalseForNonEmptySequence(t *testing.T) {
	t.Parallel()
	s := sequence.Sequence[string]{"a", "b", "c"}
	got := s.Empty()
	if got {
		t.Fatal("true for non-empty sequence")
	}
}

// Generic types can be instantiated with interface types
func TestInstantiation(t *testing.T) {
	x := sequence.Sequence[error]{errors.New("oh no")}

	fmt.Println(x)
}

func TestMaps(t *testing.T) {
	m := sequence.Catalog[int]{}

	fmt.Println(m["bob"])

	m2 := sequence.Index[string, int]{}
	m2["bob"] = 31
	fmt.Println(m2["bob"])
}

func TestStructs(t *testing.T) {
	x := sequence.NamedThing[int]{Thing: 2, Name: "bob"}
	y := sequence.NamedThing[sequence.NamedThing[int]]{
		Thing: sequence.NamedThing[int]{
			Thing: 2,
			Name:  "Alice",
		},
		Name: "bob",
	}

	fmt.Println(x)
	fmt.Println(y)
}

func TestChannels(t *testing.T) {
	ch := make(sequence.MyChan[string])

	go func() {
		defer close(ch)
		ch <- "hello world"
	}()

	fmt.Println(<-ch)
}
