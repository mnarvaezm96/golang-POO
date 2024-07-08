package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmpolyee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmpolyee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 5
	ftEmployee.id = 2
	fmt.Printf("%v\n", ftEmployee)
}
