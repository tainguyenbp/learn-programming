package main

import "fmt"

func main() {
	if var1 := 100; var1 > 20 {
		fmt.Println("Value inside if:", var1)
	}
	// Cannot use the variable var1 here

}
