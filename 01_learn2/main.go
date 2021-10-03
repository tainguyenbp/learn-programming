package main
import "fmt"

func main() {

	fmt.Println(Sum(5))
}

func Sum(number int) (int) {
	sum := 0
	for i:=0; i < number; i++ {
		sum += i 
	}
	return sum
}
