package main

import "fmt"

// similitud a POO

type Employee struct {
	id       int
	name     string
	vatacion bool
}

func newEmployee(id int, name string, vatacion bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vatacion: vatacion,
	}
}

func main() {
	// 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	// 2
	e2 := Employee{
		id:       1,
		name:     "Name",
		vatacion: true,
	}
	fmt.Printf("%v\n", e2)

	// 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)

	//4
	e4 := newEmployee(1, "name 2", true)
	fmt.Printf("%v\n", *e4)
}
