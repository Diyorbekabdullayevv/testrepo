package goconcurrency

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) increment() {
	// c.mu.Lock()
	// defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func WorkWithMutexes() {
	WorkWithMutexes1()
	// WorkWithMutexes2()
}

func WorkWithMutexes1() {

	var wg sync.WaitGroup
	counter := counter{}

	numGoroutines := 10

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// counter.count++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d!", counter.getValue())

}

func WorkWithMutexes2() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)

	increament := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numGoroutines {
		go increament()
	}

	wg.Wait()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("Final counter value: %d\n", counter)
		}
	}
}
