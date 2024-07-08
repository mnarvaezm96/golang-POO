package main

// funciones variadicas: sirve para usar varios variables sin tener que poner una cantidad x de parametros

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, quad int) {
	//return 2 * x, 3 * x, 4 * x
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return

}

func main() {
	fmt.Println(sum(4, 5, 6, 7, 8, 8))
	fmt.Println(sum(1, 2, 3))

	printNames("Alince", "teo", "garro", "cualquiera")
	fmt.Print(getValues(2))
}
