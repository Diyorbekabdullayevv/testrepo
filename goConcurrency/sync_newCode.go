package goconcurrency

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewBuffer(size int) *buffer {
	b := &buffer{items: make([]int, 0, size)}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) Produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == bufferSize {
		b.cond.Wait()
	}

	b.items = append(b.items, item)
	fmt.Println("Produced:", item)
	b.cond.Signal()
}

func (b *buffer) Consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	for len(b.items) == 0 {
		b.cond.Wait()
	}

	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed:", item)
	b.cond.Signal()

	return item
}

func Producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		b.Produce(i + 100)
		time.Sleep(time.Millisecond * 100)
	}
}

func Consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		b.Consume()
		time.Sleep(time.Millisecond * 200)
	}
}

func SyncNewCode() {
	SyncNewCode1()
}

func SyncNewCode1() {
	var wg sync.WaitGroup
	buffer := NewBuffer(bufferSize)

	wg.Add(2)
	go Producer(buffer, &wg)

	go Consumer(buffer, &wg)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Items slice value:", buffer.items)
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Items slice value:", buffer.items)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Items slice value:", buffer.items)
	wg.Wait()
}
