package main

import (
	"fmt"
	"reflect"
	"time"
)

// Variables
var second_number int64
var a_float float64
var name string
var is_male bool

// Constants
const PUBLISHER = "Villaxlnc"

func main() {
	var first_number = 5
	// ^ type inferred
	fmt.Println(first_number) // 5

	fmt.Println("second_number", second_number) // 0
	fmt.Println("a_float", a_float)             // 0
	fmt.Println("name", name)                   // ""
	fmt.Println("is_male", is_male)             // false

	var cgpa float32 = 4.5
	fmt.Println(cgpa)

	// Multiple variable declaration
	username, age := "@tcitrogg", 20
	fmt.Printf("Username: %s\nAge: %d\n", username, age)

	fmt.Printf("Publisher: %s\n", PUBLISHER)

	// Raw string
	address := `
	Squad 0 Court,
	Laughtale Island
	`

	fmt.Println(address)

	fmt.Printf(`
	Username: "%s" has %d letters`, username, len(username))

	// Using %T to get the type of a value
	current_time := time.Now()
	fmt.Println("\nCurrent Time: ", current_time)
	fmt.Printf("%T\n", current_time)

	// Using `reflect` package
	fmt.Printf("%s\n", reflect.TypeOf(current_time))
	fmt.Printf("%s\n", reflect.ValueOf(current_time).Kind())

	// Type convertion
	var user_age int
	fmt.Scanf("Enter your age: %d", &user_age)
	fmt.Println("> Your age is : ", user_age)
}
