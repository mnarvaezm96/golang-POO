package main

import "fmt"

// similitud a POO

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Printf("%v", e)
	e.id = 1
	e.name = "teo"
	fmt.Printf("%v", e)
	e.SetId(5)
	e.SetName("Teo2")
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

}
