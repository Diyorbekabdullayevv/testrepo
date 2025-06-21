package advanced // context.go

import (
	"context"
	"fmt"
	"log"
	"time"
)

func CheckEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation canceled!"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even!", num)
		} else {
			return fmt.Sprintf("%d is odd!", num)
		}
	}
}

func DoWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func LogWithContext(ctx context.Context, message string) {
	requestIDVal := ctx.Value("requestID")
	log.Printf("RequestID: %v - %v", requestIDVal, message)
}

func WorkWithContect() {
	// WorkWithContect1()
	// WorkWithContect2()
	WorkWithContect3()
}

func WorkWithContect1() {

	todoContext := context.TODO()
	contextBkg := context.Background()

	ctx := context.WithValue(todoContext, "name", "Zafar")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	bkg := context.WithValue(contextBkg, "city", "Khonka")
	fmt.Println(bkg)
	fmt.Println(bkg.Value("city"))
}

func WorkWithContect2() {

	ctx := context.TODO()

	result := CheckEvenOdd(ctx, 5)
	fmt.Println("Result of context.TODO():", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	result = CheckEvenOdd(ctx, 10)
	fmt.Println("Result of context.Background():", result)

	time.Sleep(2 * time.Second)
	result = CheckEvenOdd(ctx, 15)
	fmt.Println("Result after timeout:", result)
}

func WorkWithContect3() {
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	ctx = context.WithValue(ctx, "requestID", "378endhjwue8uekf")

	go DoWork(ctx)

	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request ID found!")
	}

	LogWithContext(ctx, "This a test log message!")
}
