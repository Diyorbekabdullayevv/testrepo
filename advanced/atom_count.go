package advanced // atom_count.go

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment1() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) getValue1() int64 {
	return atomic.LoadInt64(&ac.count)
}

func WorkWithAtomicCounters() {
	var wg sync.WaitGroup
	var numGoroutines = 10
	counter := &AtomicCounter{}
	// value := 0

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment1()
				// value++
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.getValue1())
	// fmt.Println("Final counter value:", value)

}
