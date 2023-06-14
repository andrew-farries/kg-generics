package stringy

// The type set of a constraint is the set of all types that satisfy it.

type OnlyInt interface {
	int | uint32
}

type Foo interface {
	OnlyInt | string
}

// Can use type literals in constraints
// But it's pretty useless since you can't access the structs fields
// through the interface.
type Pointish interface {
	struct{ X, Y int }
}

func Double[T OnlyInt](x T) T {
	return x * 2
}

// Cant access x.X or x.Y in here
func Diag[T Pointish](x T) T {
	return struct{ X, Y int }{X: 1, Y: 1}
}
