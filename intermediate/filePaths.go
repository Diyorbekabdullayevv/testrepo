package main

// import (
// 	"fmt"
// 	"path/filepath"
// 	"strings"
// )

// func main() {

// 	relativePath := "./data/file.txt"
// 	absolutePath := "C:\\Users\\YourName\\Documents\\file.txt"

// 	joinedPaths := filepath.Join("downloads", "file.zip")
// 	fmt.Println("Joined path:", joinedPaths)
// 	fmt.Println("")

// 	normPath := filepath.Clean("./join/./clean/../whatever/../jumanji.com")
// 	fmt.Println("Normalised path:", normPath)
// 	fmt.Println("")

// 	path, file := filepath.Split("join/clean/whatever/jumanji.com")
// 	fmt.Println("Path:", path)
// 	fmt.Println("File:", file)
// 	fmt.Println("Also file:", filepath.Base("join/clean/whatever/jumanji.com"))

// 	fmt.Println("Is relativePath absolute:", filepath.IsAbs(relativePath))
// 	fmt.Println("Is absolutePath absolute:", filepath.IsAbs(absolutePath))

// 	fmt.Println("File extention:", filepath.Ext(file))
// 	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

// 	rel, err := filepath.Rel("a/b", "a/b/c/d/file")
// 	if err != nil {
// 		return
// 	}

// 	fmt.Println(rel)

// 	rel, err = filepath.Rel("a/j", "a/b/c/d/file")
// 	if err != nil {
// 		return
// 	}

// 	fmt.Println(rel)

// 	absPath, err := filepath.Abs(relativePath)
// 	if err != nil {
// 		return
// 	}

// 	fmt.Println("Absolute path:", absPath)
// }
