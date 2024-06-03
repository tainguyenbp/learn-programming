package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Please enter a number:")

	// Reading input from the user
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("You entered:", num)

	// Using a switch statement to handle different cases
	switch {
	case num < 0:
		fmt.Println("You entered a negative number.")
	case num == 0:
		fmt.Println("You entered zero.")
	case num > 0 && num <= 10:
		if num == 5 {
			fmt.Println("You entered exactly five.")
		} else {
			fmt.Println("You entered a positive number between 1 and 10.")
		}
	case num > 10 && num <= 20:
		if num%2 == 0 {
			fmt.Println("You entered an even number between 11 and 20.")
		} else {
			fmt.Println("You entered an odd number between 11 and 20.")
		}
		fallthrough
	case num > 20 && num <= 30:
		if num%5 == 0 {
			fmt.Println("You entered a multiple of 5 between 21 and 30.")
		} else {
			fmt.Println("You entered a number between 21 and 30.")
		}
		fallthrough
	case num > 30:
		if num > 50 {
			fmt.Println("You entered a number greater than 50.")
		} else {
			fmt.Println("You entered a number between 31 and 50.")
		}
	default:
		fmt.Println("This case should never be reached.")
	}

	// Example of handling multiple specific cases with nested if-else
	switch num {
	case 1, 2, 3:
		fmt.Println("You entered a small number (1, 2, or 3).")
		if num == 2 {
			fmt.Println("Two is the only even small number.")
		}
	case 10, 20, 30:
		fmt.Println("You entered a round number (10, 20, or 30).")
		if num == 20 {
			fmt.Println("Twenty is a nice round number.")
		}
	default:
		fmt.Println("You entered a number that is neither small nor round.")
	}
}
