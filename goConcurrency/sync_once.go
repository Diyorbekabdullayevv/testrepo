package goconcurrency

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Once only printed code line!")
}

func SyncOnce() {
	SyncOnce1()
}

func SyncOnce1() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine #", i)
			once.Do(initialize)
		}()
	}
	wg.Wait()
}
