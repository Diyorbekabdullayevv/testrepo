package advanced

import (
	"fmt"
	"time"
)

func WorkWithUnBufChans() {
	ch := make(chan int)
	go func() {

		time.Sleep(3 * time.Second)
		// fmt.Println(<-ch)
		fmt.Println("3 second goroutine finished!")
	}()
	ch <- 1
	// go func() {

	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("2 second goroutine finished!")
	// }()

	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("3 second goroutine finished!")
	// }()

	// reciever := <-ch
	// fmt.Println(reciever)
	fmt.Println("End of the program!")
}
