package main

import "fmt"

func main() {
	if var1 := 10; var1 > 20 {
		fmt.Println("Value inside if:", var1)
	} else {
		// Can use var1 here
		fmt.Println("Value inside else:", var1)
	}

}
