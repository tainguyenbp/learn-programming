package main

import "fmt"

func multiply(x int, y int) int {
	return x * y

}

func multiplys(x int, y int) int {
	var z int
	z = x * y
	return z
}

func addTwo(x int, y int) (int, int) {
	return x + 2, y + 2
}

func addTwos(x, y int) (int, int) {
	return x + 2, y + 2
}

func addTwo2(x int, y int) (xPlusTwo int, yPlusTwo int) {
	xPlusTwo = x + 2
	yPlusTwo = y + 2
	return xPlusTwo, yPlusTwo
}

func init() {
	fmt.Println("Executing init function!")
}

func main() {

	fmt.Println("multiply: ", multiply(10, 20))
	fmt.Println("multiplys: ", multiplys(10, 30))

	addTwo_x, addTwo_y := addTwo(5, 10)

	addTwos_x, addTwos_y := addTwos(9, 8)

	fmt.Println(addTwo(5, 10))
	fmt.Println("addTwo: ", addTwo_x, addTwo_y)

	fmt.Println(addTwos(9, 8))
	fmt.Println("addTwos: ", addTwos_x, addTwos_y)

	addTwo2_x, addTwo2_y := addTwo2(7, 7)
	fmt.Println(addTwo2(20, 30))
	fmt.Println("addTwo2: ", addTwo2_x, addTwo2_y)

	fmt.Println("Executing main!")
}
