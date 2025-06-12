package advanced

import (
	"fmt"
	"time"
)

func WorkWithBufChansAll() {
	WorkWithBufChans1()
	// WorkWithBufChans2()
}

func WorkWithBufChans1() {

	// variable := make(chan int, 2)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Recieving from buffer!")
	go func() {
		fmt.Println("2 second timer starts!")
		time.Sleep(2 * time.Second)
		fmt.Println("Recieved:", <-ch)
	}()
	// fmt.Println("Blocking stars here!")
	ch <- 3
	// fmt.Println("Blocking ends here!")
	// fmt.Println("Recieved:", <-ch)
	// fmt.Println("Recieved:", <-ch)
}

func WorkWithBufChans2() {
	ch := make(chan int, 2)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		ch <- 2

	}()
	fmt.Println("Value:", <-ch)
	fmt.Println("Value:", <-ch)
	fmt.Println("The end!")
}
