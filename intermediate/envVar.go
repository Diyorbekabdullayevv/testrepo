package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {

// 	// user := os.Getenv("USER")
// 	// home := os.Getenv("HOME")

// 	// fmt.Println("User env var:", user)
// 	// fmt.Println("Home env var:", home)

// 	// err := os.Setenv("FRUIT", "APPLE")
// 	// if err != nil {
// 	// 	fmt.Println("Environment vareable setting failed:", err)
// 	// }

// 	// fmt.Println("Fruit env var:", os.Getenv("FRUIT"))

// 	// // for _, e := range os.Environ() {
// 	// // 	kvpair := strings.SplitN(e, "=", 2)
// 	// // 	fmt.Println(kvpair[0])
// 	// // }

// 	// err = os.Unsetenv("FRUIT")
// 	// if err != nil {
// 	// 	fmt.Println("Environment vareable unsetting failed:", err)
// 	// }

// 	// fmt.Println("Fruit env var:", os.Getenv("FRUIT"))

// 	// fmt.Println("-------------------------------------------------------------")

// 	// str := "a=b=c=d=e"

// 	// fmt.Println(strings.SplitN(str, "=", -1))
// 	// fmt.Println(strings.SplitN(str, "=", 0))
// 	// fmt.Println(strings.SplitN(str, "=", 1))
// 	// fmt.Println(strings.SplitN(str, "=", 2))
// 	// fmt.Println(strings.SplitN(str, "=", 3))
// 	// fmt.Println(strings.SplitN(str, "=", 4))

// 	// fmt.Println("-------------------------------------------------------------")

// 	// fmt.Println(strings.Split("Mexanizatsiyalashtirilganlardanmisizlar", "a"))
// }

// var builder strings.Builder
// 	var userInput int

// 	for {
// 		randNum1 := rand.Intn(50)
// 		randNum2 := rand.Intn(50)
// 		var num int = randNum2 + randNum1

// 		fmt.Printf("What is %d + %d?\nYour answer: ", randNum1, randNum2)
// 		fmt.Scan(&userInput)

// 		builder.Reset()
// 		builder.WriteString("You")
// 		builder.WriteRune(' ')
// 		builder.WriteString("are")

// 		if userInput == num {
// 			builder.WriteRune(' ')
// 			builder.WriteString("not")
// 		}
// 		builder.WriteRune(' ')
// 		builder.WriteString("stupid!")
		
// 		if userInput == -1 {
// 			fmt.Println("Nice job. Come again some time!")
// 			break
// 		}

// 		fmt.Println(builder.String())
// 	}
