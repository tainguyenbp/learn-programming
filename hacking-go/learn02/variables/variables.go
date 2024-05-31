package main

import "fmt"

func main() {
	var a, b int = 10, 20

	var d int = 100
	var f int = 200

	var sampleInt, sampleBoolean, sampleString = 50, true,
		"Hello tainguyenbp"

	var sampleInt_single = 50
	var sampleBoolean_single = false
	var sampleString_single = "Hello tainguyen"

	fmt.Println("value variable is: ", a, b, d, f, sampleInt, sampleBoolean, sampleString)
	fmt.Println(a, b, d, f, sampleInt, sampleBoolean, sampleString)

	fmt.Println("value variable single is: ", sampleInt_single, sampleBoolean_single, sampleString_single)
	fmt.Println(sampleInt_single, sampleBoolean_single, sampleString_single)

}
