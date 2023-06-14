package product

import "golang.org/x/exp/constraints"

type Numeric interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Product[T Numeric](x, y T) T {
	return x * y
}

func AddAnything[T constraints.Integer](x, y T) T {
	return x + y
}

func Greater[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func Identity[T any, U any](x T, y U) (T, U) { return x, y }

// constraints.Ordered is too restrictive here.
// Types like bool have equality defined, but no ordering.
// Also structs have equality but no ordering.
// Also complex numbers, pointers, channels, arrays, and interfaces.
//
// `comparable` is a compiler built-in; it can't be defined using an
// interface in source code.
func Equal[T comparable](x, y T) bool {
	return x == y
}
