package main

import "fmt"

func main() {
	var var1 int = 10 // Declare and initialize var1

	switch {
	case var1 > 20:
		fmt.Println("var1 is greater than 20")
	case var1 > 15:
		fmt.Println("var1 is greater than 15 but less than or equal to 20")
	case var1 > 10:
		fmt.Println("var1 is greater than 10 but less than or equal to 15")
	default:
		fmt.Println("var1 is 10 or less")
	}
}
