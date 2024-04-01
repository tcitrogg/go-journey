package main

import "fmt"

func main() {

	// Incrementing i : From 0 to 7
	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}

	// Decrementing i : From 7 to 0
	for i := 7; i >= 0; i-- {
		fmt.Println(i)
	}

	// Fibonacci
	max := 100
	a, b := 0, 1
	for b <= max {
		println(b)
		a, b = b, a+b
	}
}
