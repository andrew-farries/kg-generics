package sequence

type Sequence[E any] []E

func (s Sequence[E]) Empty() bool {
	return len(s) == 0
}

func (s Sequence[E]) First() E {
	var fst E

	if len(s) > 0 {
		fst = s[0]
	}
	return fst
}

// We can't have type parameters on methods :(
// func (s Sequence[E]) PrintWith[T any](x T) {
//   fmt.Println(x, s)
// }

// Maps

type Catalog[V any] map[string]V

type Index[K comparable, V any] map[K]V

// Structs

type NamedThing[T any] struct {
	Thing T
	Name  string
}

// Channels
type MyChan[E any] chan E
