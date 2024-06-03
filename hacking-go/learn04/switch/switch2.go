package main

import "fmt"

func main() {
	var num int
	fmt.Println("Please enter a number:")

	// Reading input from the user
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Please enter number type int, error reading input:", err)
		return
	}

	fmt.Println("You entered:", num)

	// Using switch statement to handle different cases
	switch num {
	case 1:
		fmt.Println("You entered one.")
	case 2:
		fmt.Println("You entered two.")
	default:
		fmt.Println("You entered a number other than one or two.")
	}
}
