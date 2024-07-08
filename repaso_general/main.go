package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(y)
	fmt.Println(x)
	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil {
		fmt.Printf("%\n", err)
	} else {
		fmt.Println(myValue)
	}
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index, value)
	}
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index, value)
	}
	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)

}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
