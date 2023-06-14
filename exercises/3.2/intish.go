package intish

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Intish interface {
	~int
}

func IsPositive[T Intish](v T) bool {
	return v > 0
}

// This is interesting.
// This interface is a constraint (a type set interface).
// It's type set is the intersection of the types satisfied by each line.
// 'any type whose underlying type is `int` and as a String() method'
// Pretty rare that you need this kind of interface.
// It's still a type set interface, so only can be used in type constraints.
type IntStringer interface {
	~int
	fmt.Stringer
}

type MyCoolInt int

func (m MyCoolInt) String() string {
	return fmt.Sprintf("cool int: %d", m)
}

func Foo[T IntStringer](x T) string {
	return x.String()
}

// This constraint interface defines an empty type set
type Oops interface {
	int
	string
}

// An example of an interface literal
// (it's a useless use of generics)
func Stringify[T interface{ String() string }](x T) string {
	return x.String()
}

// Can use types in interface literals too:
func Increment[T interface{ ~int }](x T) T {
	return x + 1
}

// The interface keyword can be omitted.
// `interface` can only be omitted when the constaint contains exactly one
// type element.
// Can't omit it for interfaces with methods either.
func Plus2[T ~int](x T) T {
	return x + 2
}

func Contains[T constraints.Ordered](xs []T, x T) bool {
	for _, v := range xs {
		if v == x {
			return true
		}
	}
	return false
}

type AnotherInt int

func (x AnotherInt) Equal(y AnotherInt) bool {
	return x == y
}

func ContainsFn[T interface{ Equal(T) bool }](xs []T, x T) bool {
	for _, v := range xs {
		if v.Equal(x) {
			return true
		}
	}
	return false
}

// Or with a named, generic interface in the constraint
type Equaler[T any] interface {
	Equal(T) bool
}

func ContainsFn2[T Equaler[T]](xs []T, x T) bool {
	for _, v := range xs {
		if v.Equal(x) {
			return true
		}
	}
	return false
}
