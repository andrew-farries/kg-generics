package print

import (
	"fmt"
	"io"
)

func PrintAnythingTo[T any](w io.Writer, x T) {
	fmt.Fprintln(w, x)
}
