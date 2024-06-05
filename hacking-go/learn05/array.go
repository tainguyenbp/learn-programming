package main

import (
	"fmt"
	"strconv"
)

func array_2_map_advan() {
	anArray := [4]int{1, -2, 14, 0}
	aMap := make(map[string]int)

	length := len(anArray)
	for i := 0; i < length; i++ {
		fmt.Printf("%s ", strconv.Itoa(i))
		aMap[strconv.Itoa(i)] = anArray[i]
	}
	fmt.Printf("\n")

	for key, value := range aMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func array_advan() {
	myArray := [4]int{1, 2, 4, -4}
	length := len(myArray)
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", myArray[i])
	}
	fmt.Printf("\n")

	otherArray := [...]int{0, 2, -2, 6, 7, 8}
	for _, number := range otherArray {
		fmt.Printf("%d ", number)
	}
	fmt.Printf("\n")

	twoD := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	twoD[1][2] = 15

	for _, number := range twoD {
		for _, other := range number {
			fmt.Printf("%d ", other)
		}
	}
	fmt.Printf("\n")

	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	threeD[0][1][1] = -1
	for _, number := range threeD {
		fmt.Printf("%d ", number)
	}
	fmt.Printf("\n")
}

func main() {
	var a [5]int
	a[0] = 10
	a[4] = 20
	fmt.Println(a) // [10 0 0 0 20]
	// Array can be initialized during creation
	// characters[2] is empty
	characters := [2]string{"Ender", "Pentra"}
	fmt.Println(characters) // [Ender Pentra ]
	fmt.Println("Test Array Advance: ")
	array_advan()
	fmt.Println("Test Array 2 map Advance: ")
	array_2_map_advan()
}
