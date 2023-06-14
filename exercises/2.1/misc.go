package print

import (
	"fmt"
	"sync"
)

func Identity[T any](x T) T { return x }

func Len[T any](xs []T) int {
	return len(xs)
}

func Drain[T any](ch <-chan T) {
	go func() {
		for range ch {
		}
		fmt.Println("channel drained")
	}()
}

func Merge[T any](chs ...<-chan T) <-chan T {
	out := make(chan T)
	wg := sync.WaitGroup{}
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(ch <-chan T) {
			for val := range ch {
				out <- val
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

type Bunch[T any] []T
