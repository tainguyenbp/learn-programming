package main

import "fmt"

func main() {

	// if else if else if else if else
	// trong golab khong can su dung break

	number := 3
	switch {
	case number > 10:
		fmt.Println("number > 10")
	case number == 10:
		fmt.Println("number = 10")
	default:
		fmt.Println("unknown")
	}

}
