package goconcurrency

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func SyncPool() {
	SyncPool1()
}

func SyncPool1() {

	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new person!")
			return &Person{}
		},
	}

	person1 := pool.Get().(*Person)
	person1.Name = "Samandar"
	person1.Age = 138
	fmt.Println("Got the first person:", person1)

	fmt.Printf("Person1 name: %s. Person1 age: %d\n", person1.Name, person1.Age)
	pool.Put(person1)
	fmt.Println("Returned person1 to pool!")
	fmt.Println("")

	person2 := pool.Get().(*Person)
	person2.Name = "Quramboy"
	person2.Age = 74
	fmt.Println("Got the second person:", person2)

	fmt.Printf("Person2 name: %s. Person2 age: %d\n", person2.Name, person2.Age)
	pool.Put(person2)
	fmt.Println("Returned person2 to pool!")

}
