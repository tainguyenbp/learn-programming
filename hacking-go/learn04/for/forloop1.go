package main

import "fmt"

func main() {

	// var sum int
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// // var sum int
	sum1, i1 := 0, 0
	for i1 < 30 {
		sum1 += i1
		i1++
	}

	fmt.Println(sum1)

	// // var sum int
	sum2 := 1
	for i2 := 0; i2 < 40; i2++ {
		sum2 += i2
	}
	fmt.Println(sum2)

}
