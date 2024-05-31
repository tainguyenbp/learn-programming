package main

import "fmt"

func main() {
	var (
		sampleInt     = 500
		sampleBoolean = true
		sampleString  = "Hello"
	)

	sampleInt1, sampleBoolean1, sampleString1 := 4000, true, "Hello Tai Nguyen"

	fmt.Println(sampleInt, sampleBoolean, sampleString)
	fmt.Println(sampleInt1, sampleBoolean1, sampleString1)

}
