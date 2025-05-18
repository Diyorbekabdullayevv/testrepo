package main

// import "fmt"

// func main() {

// 	// valueOfi := adder()
// 	// fmt.Println(valueOfi())
// 	// fmt.Println(valueOfi())
// 	// fmt.Println(valueOfi())
// 	// fmt.Println(valueOfi())
// 	// fmt.Println(valueOfi())

// 	// valueOfi2 := adder()
// 	// fmt.Println(valueOfi2())
// 	// fmt.Println(valueOfi2())
// 	// fmt.Println(valueOfi2())
// 	// fmt.Println(valueOfi2())
// 	// fmt.Println(valueOfi2())

// 	subtracter := func() func(int) int {
// 		var countdown int = 99
// 		return func (x int) int {
// 			countdown -=x
// 			return countdown
// 		}
// 	}()

// 	multiplier := func () int {
// 	i := 1
// 	fmt.Print("Initial subtration ", subtracter(i), ": ")
// 	i++
// 	return i
// 	}

// 	fmt.Println(multiplier())
// 	fmt.Println(multiplier())
// 	fmt.Println(multiplier())
// 	fmt.Println(multiplier())
// 	fmt.Println(multiplier())
// 	fmt.Println(multiplier())

// }

// func adder() func() int {
// 	var i int
// 	fmt.Println("Value of i is", i)
// 	return func() int {
// 		i++
// 		fmt.Print("Added 1 to i: ")
// 		return i
// 	}
// }