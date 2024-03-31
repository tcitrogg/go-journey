package main

import "fmt"

func bridgeEthAndXrp() (int, bool) {
	// Eth no Xrp desu
	return 8, false
}

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

	// bridgeValue, bridgeError := bridgeEthAndXrp()

	if bridgeValue, bridgeError := bridgeEthAndXrp(); bridgeError {
		fmt.Printf("(x) Error: %skyddcv	\n", bridgeError)
	} else {
		fmt.Println(bridgeValue)
	}

	// For weekdays
	day := 5
	dayInWeek := ""

	switch day {
	case 1:
		dayInWeek = "Sunday"
	case 2:
		dayInWeek = "Monday"
	case 3:
		dayInWeek = "Tuesday"
	case 4:
		dayInWeek = "Wednesday"
	case 5:
		dayInWeek = "Thursday"
	case 6:
		dayInWeek = "Friday"
	case 7:
		dayInWeek = "Saturday"
	default:
		dayInWeek = "--error--"
	}
	fmt.Println(dayInWeek)

	// For grading
	score := 65
	grade := ""

	switch {
	case score < 40:
		grade = "F"
	case score < 45:
		grade = "E"
	case score < 50:
		grade = "D"
	case score < 60:
		grade = "C"
	case score < 70:
		grade = "B"
	default:
		grade = "A"
	}

	fmt.Println(grade)
}
