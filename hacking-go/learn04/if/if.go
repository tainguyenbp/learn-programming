package main

import "fmt"

func main() {
	a := 20
	b := 20

	if b > a {

		fmt.Println(b, ">", a)
	}
	if b < a {
		fmt.Println(b, "<", a)
	}
	if b == a {
		fmt.Println(b, "=", a)
	}
}
