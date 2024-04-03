package main

import (
	"fmt"
	"strings"
	"time"
)

func displayDate(format string, prefix string) {
	fmt.Println(prefix, time.Now().Format(format))
}

// Function with pointers
func swap(a, b *int) {
	// Pointer of a and b
	*a, *b = *b, *a
}

// Returning values
func add_two_numbers(a, b int) int {
	return a + b
}

// Returning multiple values
func swapValues(x, y string) (string, string) {
	return y, x
}

// Multiple arguments
func add_str(sep string, texts ...string) string {
	return strings.Join(texts, sep)
}

func add_nums(numbers ...int) int {
	result := 0
	for value := range numbers {
		result += value
	}
	return result
}

// Anonymous Function, Variable function type
var i func() int

func do_some() int {
	return 8
}

// Closure
func filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	year, month, date := time.Now().Date()
	fmt.Println("The Year:", year, "- Month:", month, "- Date:", date)

	displayDate("Mon 2006-01-02 15:04:05", "Current Date and Time")

	x := 4
	y := 7
	// Addrees of x and y
	swap(&x, &y)

	fmt.Println(x, y)

	handle, firstname := "Tcitrogg", "Radiance"
	handle, firstname = swapValues(handle, firstname)
	fmt.Printf("handle: %s, firstname: %s\n", handle, firstname)

	fmt.Println(add_str(" | ", "Tcitrogg", "Bnierimi", "Tsurgeon"))

	fmt.Println(add_nums(1, 2, 3, 4, 5, 6, 7, 8, 9))

	// To use i
	fmt.Println(i())

	// Running Closure
	list := []int{1, 2, 3, 4, 5}
	evens := filter(list,
		func(i int) bool {
			return i%2 == 0 // Even numbers
			// return i > 3 // Numbers greater than 3
		})
	fmt.Println(evens)
}
