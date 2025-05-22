package main

// import (
// 	"crypto/rand"
// 	"crypto/sha256"
// 	"encoding/base64"
// 	"fmt"
// 	"io"
// )

// func main() {

// 	password := "myPasword2003"

// 	// hashedPassword256 := sha256.Sum256([]byte(password))
// 	// hashedPassword512 := sha512.Sum512([]byte(password))

// 	// fmt.Println("Password:", password)
// 	// fmt.Println("Hashed password:", hashedPassword256)
// 	// fmt.Println("Hashed password (512):", hashedPassword512)

// 	// fmt.Printf("ShA-256 Hash hex val: %x\n", hashedPassword256)
// 	// fmt.Printf("ShA-256 Hash hex val (512): %x\n", hashedPassword512)
// 	// // fmt.Println("Hashed password (string):", string(hashedPassword))

// 	salt, err := generateSalt()
// 	fmt.Println("Original salt:", salt)
// 	fmt.Printf("Original salt (hex): %x\n", salt)
// 	if err != nil {
// 		fmt.Println("Error salting:", err)
// 	}

// 	SignUpHash := hashPassword(password, salt)
// 	saltStr := base64.StdEncoding.EncodeToString(salt)
// 	fmt.Println("Hashed password:", SignUpHash)
// 	fmt.Println("Salted string:", saltStr)
// 	hashOriPassword := 	sha256.Sum256([]byte(password))
// 	fmt.Println("Just hash password:", base64.StdEncoding.EncodeToString(hashOriPassword[:]))

// 	decodeHash, err := base64.StdEncoding.DecodeString(saltStr)
// 	if err != nil {
// 		fmt.Println("Unable to decode hash:", err)
// 	}

// 	loginHash := hashPassword(password, decodeHash)

// 	if loginHash == SignUpHash {
// 		fmt.Println("Approved!")
// 	} else {
// 		fmt.Println("Rejected!")
// 	}
// }

// func generateSalt() ([]byte, error) {

// 	salt := make([]byte, 16)
// 	_, err := io.ReadFull(rand.Reader, salt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return salt, nil
// }

// func hashPassword(password string, salt []byte) string {
// 	saltedPassword := append(salt, []byte(password)...)
// 	hash := sha256.Sum256(saltedPassword)

// 	return base64.StdEncoding.EncodeToString(hash[:])
// }
