package funcmap

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type FuncMap[T, U any] map[string]func(T) U

func (m FuncMap[T, U]) Apply(f string, x T) U {
	fn, ok := m[f]
	if !ok {
		panic("function not found")
	}

	return fn(x)
}

func Reverse[T any](xs []T) {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
}

func Map[T, U any](xs []T, fn func(T) U) (out []U) {
	for _, x := range xs {
		out = append(out, fn(x))
	}
	return
}

func IsEven[T constraints.Integer](x T) bool {
	return x%2 == 0
}

func Filter[E any](xs []E, pred func(E) bool) []E {
	out := make([]E, 0, len(xs))

	for _, x := range xs {
		if pred(x) {
			out = append(out, x)
		}
	}
	return out
}

func Reduce[E, T any](xs []E, fn func(T, E) T, seed T) T {
	acc := seed
	for i := 0; i < len(xs); i++ {
		acc = fn(acc, xs[i])
	}
	return acc
}

func ConcurrentFilter[E any](xs []E, pred func(E) bool) chan E {
	ch := make(chan E)
	wg := sync.WaitGroup{}
	wg.Add(len(xs))

	for _, x := range xs {
		go func(x E) {
			defer wg.Done()

			if pred(x) {
				ch <- x
			}
		}(x)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
