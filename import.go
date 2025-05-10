package main

import "fmt"

// import "fmt"

type Calculater struct {
	Name   string
	Age    int
	UserID int
}

func calculater() *Calculater {
	calculater := Calculater{
		Name:   "",
		Age:    0,
		UserID: 0,
	}
	return &calculater
}

func main() {
	// value := test()
	// value = "Ikrom"
	// fmt.Println(value)

	// myStruct := calculater()

	// myStruct.Name = "Iskandar"
	// fmt.Println(myStruct)

	// calculateRemainder()

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i := 1; i <= len(numbers); i++ {

		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}
}
