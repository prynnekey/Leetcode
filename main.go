package main

import "fmt"

func func1(a, b int) int {
	x := a + b
	func2(x)
	return x + 10086
}

func func2(x int) {
	m := x + 1
	n := x + 7
	fmt.Printf("m: %v\n", m)
	fmt.Printf("n: %v\n", n)
}

func main() {
	a, b := 2, 5
	func1(a, b)
	c := a + b
	fmt.Printf("c: %v\n", c)
}
