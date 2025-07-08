package advanced

import (
	"fmt"
	"reflect"
	"time"
)

type Greeter struct{}

func (g Greeter) Greet(fname, lname string) string {
	return "Hello" + " " + fname + " " + lname + "!"
}

func WorkWithReflect() {
	// WorkWithReflect1()
	// WorkWithReflect2()
	// WorkWithReflect3()
	WorkWithReflect4()
}

func WorkWithReflect1() {

	x := 12
	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Println("Value:", v)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	fmt.Println("Is int:", t.Kind() == reflect.Int)
	fmt.Println("Is zero:", v.IsZero())

	fmt.Println("")
	y := 10
	v1 := reflect.ValueOf(&y).Elem()
	v2 := reflect.ValueOf(&y)
	fmt.Println("V2 type:", v2.Type())
	fmt.Println("Original value:", v1.Int())

	v1.SetInt(18)
	fmt.Println("Modified value:", v1.Int())

	fmt.Println("")
	var itf interface{} = "Hello, world!"
	v3 := reflect.ValueOf(itf)

	fmt.Println("V3 type:", v3.Type())

	if v3.Kind() == reflect.String {
		fmt.Println("String value:", v3.String())
	}
}

func WorkWithReflect2() {

	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Alice", Age: 30}
	v := reflect.ValueOf(p)

	for i := range v.NumField() {
		fmt.Printf("Field %d: %v\n", i+1, v.Field(i))
	}
	fmt.Println("")

	v1 := reflect.ValueOf(&p).Elem()

	nameField := v1.FieldByName("Name")

	if nameField.CanSet() {
		nameField.SetString("Jane")
	}

	for i := range v1.NumField() {
		fmt.Printf("Field %d: %v\n", i+1, v1.Field(i))
	}

}

func WorkWithReflect3() {

	g := Greeter{}

	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)
	var method reflect.Method
	fmt.Println("Type:", t)

	for i := range t.NumMethod() {
		method = t.Method(i)
		fmt.Printf("Method %d: %v\n", i+1, method.Name)
	}

	m := v.MethodByName(method.Name)
	results := m.Call([]reflect.Value{reflect.ValueOf("Alice"), reflect.ValueOf("Martini")})
	// fmt.Println("Result:", results)
	fmt.Println("Result:", results[0].String())

}

func WorkWithReflect4() {

	ticker := time.NewTicker(500 * time.Millisecond)

	for i := range 10 {
		<- ticker.C 
		fmt.Println("Value:", i + 1)
	}
}
