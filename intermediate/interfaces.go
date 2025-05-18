package main

// import "fmt"

// // import (
// // 	"fmt"
// // 	"math"
// // )

// // type geometry interface {
// // 	area() float64
// // 	perim() float64
// // }

// // type rect struct {
// // 	width, height float64
// // }

// // type circle struct {
// // 	radius float64
// // }

// // func (r rect) area() float64 {
// // 	return r.width * r.height
// // }

// // func (r rect) perim() float64 {
// // 	return 2 * (r.width + r.height)
// // }

// // func (c circle) area() float64 {
// // 	return math.Pi * c.radius * c.radius
// // }

// // func (c circle) perim() float64 {
// // 	return 2 * math.Pi * c.radius
// // }

// // // func (c circle) diameter() float64 {
// // // 	return 2 * c.radius
// // // }

// // func measure(g geometry) {

// // 	fmt.Println(g)
// // 	fmt.Println(g.area())
// // 	fmt.Println(g.perim())
// // }

// type cooker interface {
// 	cook() string
// }

// type pilav struct {
// 	rice, meat string
// }

// type soup struct {
// 	meat, water string
// }

// func (p pilav) cook() string {
// 	return p.rice + p.meat
// }

// func (s soup) cook() string {
// 	return s.meat + "and" + s.water
// }

// func measure(c cooker) {

// 	fmt.Println(c.cook())
// 	fmt.Println("")
// }

// func main() {

// 	p := pilav{rice: "Pilav is cooked with ", meat: "meat"}
// 	s := soup{meat: "Soup is cooked with meat ", water: " water"}

// 	measure(p)
// 	measure(s)

// 	// 	r := rect{width: 3, height: 4}
// 	// 	c := circle{radius: 5}

// 	// 	measure(r)
// 	// 	measure(c)

// }
