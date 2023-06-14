package stringy

// The type set of a constraint is the set of all types that satisfy it.
// An interface that contains types (as opposed to only methods) is called
// a 'constraint'. You can't use a constraint where you would use an ordinary
// interface.
// So Go effectively has two kinds of interface now.
// Previous versions of the proposal called the new kind 'contracts' with a
// different keyword.

type OnlyInt interface {
	int | uint32
}

type Fooer interface {
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

// This doesn't work: Can't use OnlyInt outside a type constraint.
// func Foo(x OnlyInt) {}

type Cow struct{ Moo string }
type Chicken struct{ Cluck string }

type Animal interface {
	Cow | Chicken
}

type Farm[T Animal] []T

type ApproxInt interface {
	~int
}

type MyInt int

func DoubleApprox[T ApproxInt](x T) T {
	return x * 2
}
