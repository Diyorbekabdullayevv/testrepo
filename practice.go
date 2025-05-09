package main

import (
	"fmt"
	"time"
)

var currentYear = time.Now().Year()
var currentMonth = time.Now().Month()
var currentDay = time.Now().Day()

var clientAge int

func calculateUserAge() {
	
	// fmt.Print("Input your born year: ")
	// fmt.Scan(&clientAge)

	// clientAge = currentYear - clientAge
	// fmt.Println("Your age is", clientAge)

	fmt.Printf("Year: %v\nMonth: %v\nDay: %v\n", currentYear, currentMonth, currentDay)
}

func testSomething() string {
	var testSomething string
	return testSomething
}

