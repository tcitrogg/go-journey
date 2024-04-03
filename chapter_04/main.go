package main

import (
	"fmt"
	"strings"
)

func main() {

	// Incrementing i : From 0 to 7
	fmt.Println("\nIncrementing")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("\nDecrementing")
	// Decrementing i : From 7 to 0
	for i := 3; i >= 0; i-- {
		fmt.Println(i)
	}

	// Fibonacci
	fmt.Println("\nFibonacci")
	max := 100
	a, b := 0, 1
	for b <= max {
		println(b)
		a, b = b, a+b
	}

	// Infinite loop
	for {
		fmt.Println("Enter [q]uit to exit")
		input := ""
		fmt.Print("Go> ")
		fmt.Scanln(&input)
		if strings.ToLower(input) == "quit" || strings.ToLower(input) == "q" {
			// The Break keyword
			break
		} else {
			fmt.Printf("(x) You entered %s\n", input)
		}
	}

	// The Continue keyword
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		println(i)
	}

	for ax, by := 0, 1; by <= max; ax, by = by, ax+by {
		println(by)
	}

	// const OS = ["iOS"
	// "Android"
	// "SailOS"]

	var OS [3]string
	OS[0] = "iOS"
	OS[1] = "Android"
	OS[2] = "SailOS"

	// Iterating over an array
	for value := range OS {
		println(value)
	}

	// Also for strings
	title := "Demigod"
	for index, char := range title {
		fmt.Printf("Index[%d] Letter[%c]", index, char)
		fmt.Println(" Char[", char, "]")
	}

	// Multiplication Table
	fmt.Println("\nMultiplication Table")
	tableLength := 5
	counter := 1
	for counter <= tableLength {
		for nextIndex := 1; nextIndex <= tableLength; nextIndex++ {
			fmt.Printf("%d * %d = %d\n", counter, nextIndex, counter*nextIndex)
		}
		counter++
		println("------------")
	}
}
