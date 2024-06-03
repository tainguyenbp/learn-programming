package main

import "fmt"

func main() {
	var var1 int = 10 // Declare and initialize var1

	if var1 > 20 {
		fmt.Println("var1 is greater than 20")
	} else {
		if var1 > 15 {
			fmt.Println("var1 is greater than 15 but less than or equal to 20")
		} else {
			if var1 > 10 {
				fmt.Println("var1 is greater than 10 but less than or equal to 15")
			} else {
				fmt.Println("var1 is 10 or less")
			}
		}
	}
}
