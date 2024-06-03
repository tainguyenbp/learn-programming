package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = 25

	if var1 > 20 && var2 > 30 {
		fmt.Println("var1 is greater than 20 and var2 is greater than 30")
	} else if var1 > 15 || var2 > 20 {
		fmt.Println("Either var1 is greater than 15 or var2 is greater than 20")
	} else if var1 == 10 {
		fmt.Println("var1 is exactly 10")
	} else {
		fmt.Println("None of the conditions were met")
	}
}
