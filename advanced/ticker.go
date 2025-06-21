package advanced // ticker.go

import (
	"fmt"
	"time"
)

func PeriodicTask() {
	fmt.Println("Performing a periodic task at:", time.Now())
}

func WorkWithTickers() {
	// WorkWithTickers1()
	// WorkWithTickers2()
	WorkWithTickers3()
}

func WorkWithTickers1() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	// for tick := range ticker.C {
	//   fmt.Println("Tick at:", tick)
	// }
	var i int = 1
	for range 5 {
		i *= 2
		fmt.Println(i)
	}

	// for tick := range ticker.C {
	//   fmt.Println("Tick:", tick)
	// }
}

func WorkWithTickers2() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			PeriodicTask()
		}
	}
}

func WorkWithTickers3() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at:", tick)
		case <-stop:
			fmt.Println("Stopping ticker!")
			return
		}
	}
}
