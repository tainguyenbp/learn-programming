package main

import "fmt"

func main() {

	// if else if else if else if else
	// trong golab khong can su dung break

	number := 15
	switch number {
	case 1, 11, 15, 100: // hỗ trợ bắt nhiều trường hợp trong 1 case
		if number == 1 {
			fmt.Println("number = 1")
		} else if number == 11 {
			fmt.Println("number = 11")
		} else if number == 15 {
			fmt.Println("number = 15")
		} else if number == 100 {
			fmt.Println("number = 100")
		}
	case 2:
		fmt.Println("number = 2")
	case 3:
		fmt.Println("number = 3")
	case 4:
		fmt.Println("number = 4")
	case 5:
		fmt.Println("number = 5")
	case 10:
		fmt.Println("number = 10")
	default:
		fmt.Println("unknown")
	}

}
