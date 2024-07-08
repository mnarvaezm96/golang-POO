package main

// Funciones anonimas
func main() {

	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.println(y)
}
