package stringy

import (
	"fmt"
	"io"
)

// A useless use of generics
// (it doesn't relate two or more types in the signature)
func StringifyTo[T fmt.Stringer](w io.Writer, x T) {
	fmt.Fprintln(w, x)
}
