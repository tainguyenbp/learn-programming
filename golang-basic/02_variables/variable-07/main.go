package main

import "fmt"

func main() {
	// khai bao nhieu bien
	// 2. khai bao nhieu bien nhung khac kieu du lieu
	var (
		name    string
		address string
		age     int
		year    int
	)

	name = "Nguyễn Ngọc Tài"
	address = "Ho Chi Minh City, Viet Nam"
	age = 18
	year = 2000

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(age)
	fmt.Println(year)
}
