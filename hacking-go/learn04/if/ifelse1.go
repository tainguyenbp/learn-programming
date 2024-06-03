package main

import "fmt"

func main() {
	var1 := 3
	if var1 > 20 {
		fmt.Println("Value inside if:", var1)
	} else {
		if var1 > 5 {
			fmt.Println("Value inside nested if:", var1) // This block executes if var1 <= 20 and var1 > 5
		} else {
			fmt.Println("Value inside nested else:", var1) // This block executes if var1 <= 5
		}
	}

}
