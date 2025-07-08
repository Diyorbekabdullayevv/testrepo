package goconcurrency

import (
	"fmt"
	"sync"
	"time"
)

func Deadlocks() {
	Deadlocks1()
}

func Deadlocks1() {

	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
		fmt.Println("Goroutine 1 completed!")
	}()

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 2 locked mu2")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 2 locked mu1")
		mu2.Unlock()
		mu1.Unlock()
	fmt.Println("Goroutine 2 completed!")

	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed!")

	// select{}
}
