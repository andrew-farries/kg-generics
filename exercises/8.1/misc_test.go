package billable_test

import (
	"billable"
	"fmt"
	"sync"
	"testing"
)

func TestSet(t *testing.T) {
	s := billable.NewSet(1, 2, 3)

	s.Add(4, 5, 6)

	fmt.Println(s)
}

func TestSetConcurrency(t *testing.T) {
	s := billable.NewSet[int]()
	wg := sync.WaitGroup{}

	wg.Add(2000)
	for i := 0; i < 1000; i++ {
		go func(x int) {
			defer wg.Done()
			s.Add(x)
		}(i)

		go func() {
			defer wg.Done()
			_ = s.Members()
		}()
	}

	wg.Wait()

	fmt.Println(len(s.Members()))
}
