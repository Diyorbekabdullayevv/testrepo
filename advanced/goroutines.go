package advanced

import (
	"fmt"
	"time"
)

func WorkWithGoroutines() {
	var err error
	fmt.Println("This is before sleep!")
	go sleep()
	fmt.Println("This is after sleep!")
	go func() {
		err = doWork()
	}()

	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully!")
	}
}

func sleep() {

	time.Sleep(1 * time.Second)
	fmt.Println("Woke up from sleep!")

}

func printNumbers() {

	for i := 0; i < 5; i++ {
		fmt.Printf("Number: %d. Execution time: %v\n", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {

	for _, letter := range "abcde" {
		fmt.Printf("Letter: %s. Execution time: %v\n", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("something went wrong")
}
