package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strings"
// )

// func main() {

// 	// fmt.Println("He said, \"I am great\"")
// 	// fmt.Println(`He said, "I am great"`)

// 	// Compile a regex pattern to match an email address
// 	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

// 	// Test strings
// 	var (
// 		email1 string = "user@email.com"
// 		email2 string = "user2@invalid_emailOOO"
// 	)

// 	//Match
// 	fmt.Println("Email1:", re.MatchString(email1))
// 	fmt.Println("Email2:", re.MatchString(email2))

// 	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

// 	date := "2025-05-18"

// 	submatches := re.FindStringSubmatch(date)

// 	fmt.Println(submatches)
// 	fmt.Println(submatches[0])
// 	fmt.Println(submatches[1])
// 	fmt.Println(submatches[2])
// 	fmt.Println(submatches[3])
// 	invalidOutput, err := fmt.Println(submatches[3])
// 	if err != nil {
// 		panic("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa: ")
// 	}
// 	fmt.Println(invalidOutput)

	
// 	// Using "regexp"
// 	name := "Abdusattorov!"
// 	re = regexp.MustCompile(`[Aaouei]`)

// 	result := re.ReplaceAllString(name, "*")
// 	fmt.Println(result)

// 	// Using "strings"
// 	vowels := []string{"a", "e", "u", "o", "i", "A"}

// 	for _, v := range vowels{
// 			name = strings.ReplaceAll(name, v, "*")
// 	}

// 	fmt.Println(name)

// 	game := "Abdusattorov!"
// 	re = regexp.MustCompile(`(?i)Go`)

// 	fmt.Println(re.MatchString(game))

// }
