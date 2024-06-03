package main

import (
	"fmt"
	"math/rand" // This is not cryptographically secure!
	"time"
)

func main() {
	// Seeding rand
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Choosing a random number:")
	num := rand.Intn(100) // Generate a random number between 0 and 99
	fmt.Printf("Random number is: %d\n", num)

	// Using switch with advanced conditions and nested if-else
	switch {
	case num < 20:
		fmt.Println("Less than 20")
		if num < 10 {
			fmt.Println("Less than 10")
		} else {
			fmt.Println("Between 10 and 19")
		}
	case num >= 20 && num < 50:
		fmt.Println("Between 20 and 49")
		if num%2 == 0 {
			fmt.Println("Even number")
		} else {
			fmt.Println("Odd number")
		}
	case num >= 50 && num < 80:
		fmt.Println("Between 50 and 79")
		if num%5 == 0 {
			fmt.Println("Multiple of 5")
		} else {
			fmt.Println("Not a multiple of 5")
		}
	default:
		fmt.Println("80 or more")
		if num > 90 {
			fmt.Println("Greater than 90")
		} else {
			fmt.Println("Between 80 and 89")
		}
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
