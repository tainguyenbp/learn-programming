package main

import "fmt"

const Whatever = "whatever"
const helloworld = "Hello World"

const (
	Const1 = "Constant String"
	Int1   = 12345
	True   = true
)

func main() {
	fmt.Println(Whatever)
	const One = 1
	fmt.Println(One)
	fmt.Println(helloworld)
	fmt.Println(Const1, Int1, True)
}
