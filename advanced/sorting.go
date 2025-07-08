package advanced // sorting.go

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type PersonSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (ps *PersonSorter) Len() int {
	return len(ps.people)
}

func (ps *PersonSorter) Less(i, j int) bool {
	return ps.by(&ps.people[i], &ps.people[j])
}

func (ps *PersonSorter) Swap(i, j int) {
	ps.people[i], ps.people[j] = ps.people[j], ps.people[i]
}

func (by By) Sort(people []Person) {
	ps := &PersonSorter{
		people: people,
		by: by,
	}

	sort.Sort(ps)
}


func WorkWithSorting() {
	// WorkWithSorting1()
	WorkWithSorting2()
}

func WorkWithSorting1() {

	numbers := []int{5, 3, 2, 4, 1}
	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)

	strSlice := []string{"John", "Antony", "Steve", "Victor", "Walter"}
	sort.Strings(strSlice)
	fmt.Println("Sorted string:", strSlice)
}

func WorkWithSorting2() {

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}
	// fmt.Println("Unsorted by age:", people)
	// sort.Sort(ByAge(people))
	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(age).Sort(people)
	fmt.Println("Sorted by age:", people)

	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	By(name).Sort(people)
	fmt.Println("Sorted by name:", people)
}


// type ByAge []Person
// type ByName []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// func (a ByName) Len() int {
// 	return len(a)
// }

// func (a ByName) Less(i, j int) bool {
// 	return a[i].Name < a[j].Name
// }

// func (a ByName) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }
