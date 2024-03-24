package main

import "fmt"

func main() {
	num := 8

	// Comparison Operators
	// Equals to
	fmt.Println(num == 7)

	// Not equals to
	fmt.Println(num != 7)

	// Less than
	fmt.Println(num < 7)

	// Less than or equal to
	fmt.Println(num <= 7)

	// Greater than
	fmt.Println(num > 7)

	// Greater than or equal to
	fmt.Println(num >= 7)

	// Logical Operators
	// AND
	result := num > 7 && num < 10
	fmt.Println(result)

	// OR
	result = num < 10 || num > 7
	fmt.Println(result)

	// NOT
	result = !(num < 10 || num > 7)
	fmt.Println(result)

	// Conditional Statements
	num2 := 9
	if num2%2 == 1 {
		fmt.Println("(+) Odd Number")
	}

	// w/ Else block
	if num2%2 == 0 {
		fmt.Println("(+) Even Number")
	} else {
		fmt.Println("(+) Odd Number")
	}

}
