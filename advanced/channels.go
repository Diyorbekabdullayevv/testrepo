package advanced

import (
	"fmt"
	"time"
)

func WorkWithChannels() {
	// variable := make(chan type)
	greeting := make(chan string)
	greetingString := "Hello!"

	go func() {
		greeting <- greetingString
		greeting <- "World!"
		for _, e := range "abcde" {
			greeting <- "Alphabet:" + string(e)
		}
	}()

	// go func() {
	// reciever := <-greeting
	// fmt.Println(reciever)
	// reciever = <-greeting
	// fmt.Println(reciever)
	// }()

	reciever := <-greeting
	fmt.Println(reciever)
	reciever = <-greeting
	fmt.Println(reciever)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of program!")
}
