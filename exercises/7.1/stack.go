package stack

import (
	"fmt"
)

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

// Sets

type Set[E comparable] map[E]bool

func NewSet[E comparable](xs ...E) Set[E] {
	m := map[E]bool{}

	for _, x := range xs {
		m[x] = true
	}
	return m
}

func (s Set[E]) Add(xs ...E) {
	for _, x := range xs {
		s[x] = true
	}
}

func (s Set[E]) Contains(x E) bool {
	return s[x]
}

func (s Set[E]) Members() []E {
	keys := make([]E, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s Set[E]) String() string {
	return fmt.Sprintf("%v", s.Members())
}

func (s Set[E]) Union(s2 Set[E]) Set[E] {
	result := NewSet(s.Members()...)
	result.Add(s2.Members()...)

	return result
}

func (s Set[E]) Intersection(s2 Set[E]) Set[E] {
	result := NewSet[E]()

	for k := range s {
		if s2.Contains(k) {
			result.Add(k)
		}
	}

	return result
}
