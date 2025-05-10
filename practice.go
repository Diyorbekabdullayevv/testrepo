package main

import (
	"fmt"
	"time"
)

const (
	year         = 2025
	month string = "May"
	day          = 9
)

var currentYear = time.Now().Year()
var currentMonth = time.Now().Month()
var currentDay = time.Now().Day()

var remainder int = 13

var clientAge int

func calculateUserAge() {

	// fmt.Print("Input your born year: ")
	// fmt.Scan(&clientAge)

	// clientAge = currentYear - clientAge
	// fmt.Println("Your age is", clientAge)

	fmt.Printf("Year: %v\nMonth: %v\nDay: %v\n", currentYear, currentMonth, currentDay)
	fmt.Printf("Year: %v\nMonth: %v\nDay: %v\n", year, month, day)

	remainder = remainder % 7
	fmt.Println(remainder)

}

func testSomething() string {
	testSomething := ""
	return testSomething
}

func calculateRemainder() {

	var userInput string

	fmt.Print("Input number: ")
	fmt.Scan(&userInput)

	var result string

	for _, v := range userInput {

          digit := v - '0' // convert rune to digit
		if digit%2 == 0 {
			result += string(v)
		}
	}
	fmt.Println(result)
}