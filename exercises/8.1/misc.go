package billable

import (
	"sync"

	"golang.org/x/exp/maps"
)

// A concurrency safe Set type

type Set[T comparable] struct {
	mu *sync.RWMutex
	m  map[T]struct{}
}

func NewSet[T comparable](xs ...T) Set[T] {
	s := Set[T]{
		m:  make(map[T]struct{}),
		mu: &sync.RWMutex{},
	}

	for _, x := range xs {
		s.m[x] = struct{}{}
	}

	return s
}

func (s *Set[T]) Add(xs ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, x := range xs {
		s.m[x] = struct{}{}
	}
}

func (s *Set[T]) Members() (out []T) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return maps.Keys(s.m)
}
