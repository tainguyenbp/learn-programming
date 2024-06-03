package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seeding rand with the current time in nanoseconds to get different results each time
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Choosing a random number:")

	// Generating a random number between 0 and 2 (inclusive)
	switch num := rand.Intn(3); num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("3") // This will handle the case when num is 0
	}
}
