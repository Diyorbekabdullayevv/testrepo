package goconcurrency

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func PrintNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 5 {
		fmt.Println("Number:", i+1)
		time.Sleep(500 * time.Millisecond)
	}
}

func PrintLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, letter := range "ABCDE" {
		fmt.Printf("Letter: %c\n", letter)
		time.Sleep(500 * time.Millisecond)
	}
}

func HeavyTask(id int, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("Task %d is starting!\n", id)

	for range 100_000_000 {
	}
	fmt.Println(time.Now())
	fmt.Printf("Task %d is finished!\n", id)
}

func GoConcurrency() {

	// GoConcurrency1()
	GoConcurrency2()

}

func GoConcurrency1() {

	var wg sync.WaitGroup
	wg.Add(2)
	go PrintNumbers(&wg)
	time.Sleep(time.Millisecond)
	go PrintLetters(&wg)
	// time.Sleep(3 * time.Second)
	wg.Wait()
}

func GoConcurrency2() {

	var wg sync.WaitGroup
	var numThreads = 4
	runtime.GOMAXPROCS(numThreads)

	for i := range numThreads {
		wg.Add(1)
		go HeavyTask(i, &wg)
		time.Sleep(time.Millisecond)
	}

	wg.Wait()
}
