package stack

type Stack[E any] struct {
	xs []E
}

func (s *Stack[E]) Push(xs ...E) {
	s.xs = append(s.xs, xs...)
}

func (s *Stack[E]) Pop() (v E, ok bool) {
	if s.Len() == 0 {
		return v, false
	}
	v = s.xs[len(s.xs)-1]
	s.xs = s.xs[:len(s.xs)-1]
	return v, true
}

func (s *Stack[E]) Len() int {
	return len(s.xs)
}

