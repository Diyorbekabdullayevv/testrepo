package advanced

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WorkWithSignals() {
	WorkWithSignals1()
}

func WorkWithSignals1() {
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		sig := <-sigs

		switch sig {
		case syscall.SIGINT:
			fmt.Println("Received SIGINT (Interrupt)")
		case syscall.SIGTERM:
			fmt.Println("Received SIGTERM (Terminate)")
		case syscall.SIGHUP:
			fmt.Println("Received SIGHUP (Hangup)")
		}
		fmt.Println("Graceful exit!")
		os.Exit(0)
	}()

	fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}
}


