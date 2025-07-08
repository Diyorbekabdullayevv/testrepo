package goconcurrency

import (
	"fmt"
	"time"
)

func UseSelectWithFor() {
	ticker := time.NewTicker(200 * time.Millisecond)
	quit := make(chan string)

	go func() {
		time.Sleep(time.Second)
		close(quit)
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-quit:
			fmt.Println("Quitting...")
			return
		}
	}
}
