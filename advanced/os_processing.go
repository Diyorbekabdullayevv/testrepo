package advanced

import (
	"fmt"
	"os/exec"
	"strings"
)

func WorkWithProcessSpawning() {
	// WorkWithProcessSpawning1()
	// WorkWithProcessSpawning2()
	WorkWithProcessSpawning3()

}

func WorkWithProcessSpawning1() {
	cmd := exec.Command("cmd", "/C", "echo", "Hello, world!")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func WorkWithProcessSpawning2() {
	cmd := exec.Command("cmd", "/C", "grep", "foo")

	cmd.Stdin = strings.NewReader("food is bad\nbar\nbaz\n")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func WorkWithProcessSpawning3() {
	cmd := exec.Command("cmd", "/C", "sleep", "5")

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting:", err)
		return
	}
	fmt.Println("Process completed!")
}

