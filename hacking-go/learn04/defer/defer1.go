package main

import "fmt"

func defer1_advance() {
	num := 100

	// First deferred function capturing the value of num at this point
	defer func(n int) {
		fmt.Println("Deferred 1: After main returns", n) // This will print "Deferred 1: After main returns 110" after main() returns.
	}(num)

	num++ // num is incremented to 101

	// Second deferred function capturing the updated value of num after the first increment
	defer func(n int) {
		fmt.Println("Deferred 2: After main returns", n) // This will print "Deferred 2: After main returns 101" after main() returns.
	}(num)

	num += 9 // num is now 110

	fmt.Println("Inside main", num) // Prints "Inside main 110"
}

func defer1() {

	num := 100
	defer fmt.Println("After main returns", num)

	num++
	num += 9
	fmt.Println("Inside main", num)

}

func main() {
	defer1()
	defer1_advance()

}
