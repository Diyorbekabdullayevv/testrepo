package goconcurrency

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwmu     sync.RWMutex
	counter1 int
)

func ReadCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock()
	fmt.Println("Read counter:", counter1)
	rwmu.RUnlock()
}

func WriteCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	rwmu.Lock()
	counter1 = value
	fmt.Printf("Written value %d for counter!\n", value)
	rwmu.Unlock()
}

func RWMutex() {
	RWMutex1()
}

func RWMutex1() {

	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go ReadCounter(&wg)
	}

	wg.Add(1)
	time.Sleep(time.Second)
	go WriteCounter(&wg, 18)

	wg.Wait()

}
