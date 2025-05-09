package main

import "fmt"

type Calculater struct{
	Name string
	Age int
	UserID int
}

func calculater() *Calculater {
	calculater := Calculater{
		Name: "",
		Age: 0,
		UserID: 0,
	}
	return &calculater
}

func main() {
	value := test()
	value = "Ikrom"
	fmt.Println(value)

	myStruct := calculater()

	myStruct.Name = "Iskandar"
	fmt.Println(myStruct)
}
