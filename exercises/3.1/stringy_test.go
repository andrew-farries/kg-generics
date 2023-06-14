package stringy_test

import (
	"bytes"
	"fmt"
	"stringy"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type greeting struct{}

func (greeting) String() string {
	return "Howdy!"
}

func TestStringifyTo_PrintsResultOfStringMethodToSuppliedWriter(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	stringy.StringifyTo[greeting](buf, greeting{})
	want := "Howdy!\n"
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDouble(t *testing.T) {
	fmt.Println(stringy.Double(2))
}

func TestFarm(t *testing.T) {
	var f stringy.Farm[stringy.Cow] = []stringy.Cow{{Moo: "moo"}}

	// but you can't do
	// Can't use Animal outside a type constraint.
	// (it's not a type, it's a contract - the new kind of Go interface)
	// (only 'classic' interfaces can be used this way)
	// var g stringy.Farm[stringy.Animal] = []stringy.Cow{{Moo: "moo"}}

	fmt.Println(f)
}

func TestDoubleApprox(t *testing.T) {
	x := stringy.MyInt(4)
	fmt.Println(stringy.DoubleApprox(x))
}
