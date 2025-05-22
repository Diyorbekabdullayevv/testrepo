package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {

// 	file, err := os.Open("example.txt")

// 	if err != nil {
// 		fmt.Println("Failed to open file:", err)
// 		return
// 	}

// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	var keyword string = "important"
// 	lineNum := 1
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		if strings.Contains(line, keyword) {

// 			updatedLine := strings.ReplaceAll(line, keyword, "shit")
// 			fmt.Printf("%d. Filtered line: %s\n", lineNum, line)
// 			fmt.Printf("%d. Updated line: %s\n", lineNum, updatedLine)
// 			lineNum++
// 		}

// 	}

// 	err = scanner.Err()
// 	if err != nil {
// 		fmt.Println("Failed to scan file:", err)
// 		return
// 	}
// }
