package main

// import (
// 	"encoding/base64"
// 	"fmt"
// )

// func main() {

// 	data := []byte("Hello, world!")

// 	// Encoding to base64
// 	encoded := base64.StdEncoding.EncodeToString(data)
// 	fmt.Println("Encoded string:", encoded)
// 	fmt.Println("")

// 	// Decoding from base64
// 	decoded, err := base64.StdEncoding.DecodeString(encoded)
// 	if err != nil {
// 		fmt.Println("Decoding failure:", err)
// 	}

// 	fmt.Println("Decoded bytesize string:", decoded)
// 	fmt.Println("")

// 	// Decoded string itself
// 	fmt.Println("Decoded string itself:", string(decoded))

// 	// Getting lenght and ASCII codes
// 	fmt.Println(len(data))

// 	for _, v := range data {
// 		fmt.Printf("ASCII code for %c is %d\n", v, v)
// 	}

// 	// URL safe encoding
// 	data2 := []byte("Hey, hey goo~d looking")
// 	URLSafeEncoded := base64.URLEncoding.EncodeToString(data2)
// 	fmt.Println("Encoded string:", URLSafeEncoded)

// }
